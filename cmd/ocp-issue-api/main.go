package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-issue-api/internal/api"
	desc "github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
	httpPort = ":7000"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "0.0.0.0"+grpcPort, "gRPC server endpoint")
	httpEndpoint = flag.String("http-server-endpoint", "0.0.0.0"+httpPort, "HTTP server endpoint")
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

	fmt.Printf("HTTP server listening on %s\n", *httpEndpoint)
	return http.ListenAndServe(*httpEndpoint, mux)
}

func runGrpc() error {
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		return err
	}

	server := grpc.NewServer()
	desc.RegisterOcpIssueApiServer(server, api.New())

	fmt.Printf("gRPC server listening on %s\n", *grpcEndpoint)

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	go func() {
		err := runGrpc()

		if err != nil {
			log.Fatal().Msg(err.Error())
			return
		}
	}()

	if err := runHttp(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
