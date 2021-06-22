package tracing

import (
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitTracing() error {
	cfg, err := jaegercfg.FromEnv()

	if err != nil {
		return err
	}

	tracer, _, err := cfg.NewTracer()

	if err != nil {
		return err
	}

	opentracing.SetGlobalTracer(tracer)

	return nil
}
