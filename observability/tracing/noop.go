package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type NoopProvider struct {
	Provider
}

func NewNoopProvider() NoopProvider {
	return NoopProvider{}
}

func (n NoopProvider) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	return noop.Tracer{}
}

func (n NoopProvider) Shutdown(ctx context.Context) error {
	return nil
}
