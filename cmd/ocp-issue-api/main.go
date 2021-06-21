package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/ozoncp/ocp-issue-api/internal/api"
	"github.com/ozoncp/ocp-issue-api/internal/events"
	"github.com/ozoncp/ocp-issue-api/internal/flusher"
	"github.com/ozoncp/ocp-issue-api/internal/metrics"
	"github.com/ozoncp/ocp-issue-api/internal/repo"
	"github.com/ozoncp/ocp-issue-api/internal/tracing"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"github.com/pressly/goose"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

const (
	grpcPort    = ":7002"
	httpPort    = ":7000"
	metricsPort = ":9100"

	issuesChunkSize         = 2
	eventsChannelBufferSize = 50
)

var (
	grpcEndpoint    = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
	httpEndpoint    = flag.String("http-server-endpoint", "0.0.0.0"+httpPort, "HTTP server endpoint")
	metricsEndpoint = flag.String("metrics-server-endpoint", "0.0.0.0"+metricsPort, "Metrics server endpoint")
	migrationsDir   = flag.String("migrations-dir", "migrations", "directory with migration files")
)

func runHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpIssueApiHandlerFromEndpoint(ctx, gwmux, *grpcEndpoint, opts)

	if err != nil {
		return err
	}

	mux.Handle("/", gwmux)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	log.Info().Msgf("HTTP server listening on %s", *httpEndpoint)

	return http.ListenAndServe(*httpEndpoint, mux)
}

func runGrpc(repo repo.Repo) error {
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		return err
	}

	server := grpc.NewServer()
	desc.RegisterOcpIssueApiServer(server, api.New(repo, flusher.New(repo, issuesChunkSize)))

	log.Info().Msgf("gRPC server listening on %s", *grpcEndpoint)

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func runEventProducer(eventCh chan events.IssueEvent) error {
	brokers := os.Getenv("OCP_ISSUE_API_KAFKA_BROKERS")
	topic := os.Getenv("OCP_ISSUE_API_KAFKA_EVENTS_TOPIC")

	producer, err := events.NewProducer(brokers, topic)

	if err != nil {
		return err
	}

	defer producer.Close()

	for {
		select {
		case event := <-eventCh:
			err = producer.Produce(event)

			if err != nil {
				return err
			}
		}
	}
}

func runMetrics() error {
	metrics.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(*metricsEndpoint, nil)
}

func main() {
	if err := tracing.InitTracing(); err != nil {
		log.Error().Msgf("failed to init tracing: %v", err)
	}

	flag.Parse()

	db, err := sql.Open("pgx", os.Getenv("OCP_ISSUE_API_DATA_SOURCE"))

	if err != nil {
		log.Fatal().Msgf("failed to load driver: %v", err)
		return
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal().Msgf("failed to close db connection: %v", err)
		}
	}()

	if err = goose.Run("up", db, *migrationsDir); err != nil {
		log.Fatal().Msgf("failed to apply migrations: %v", err)
		return
	}

	eventCh := make(chan events.IssueEvent, eventsChannelBufferSize)

	go func() {
		if err = runEventProducer(eventCh); err != nil {
			log.Fatal().Msgf("failed to produce events: %v", err)
			return
		}
	}()

	go func() {
		if err = runGrpc(repo.New(db, eventCh)); err != nil {
			log.Fatal().Msgf("failed to run gRPC server: %v", err)
			return
		}
	}()

	go func() {
		if err = runMetrics(); err != nil {
			log.Fatal().Msgf("failed to run metrics: %v", err)
			return
		}
	}()

	if err = runHttp(); err != nil {
		log.Fatal().Msgf("failed to run HTTP server: %v", err)
	}
}
