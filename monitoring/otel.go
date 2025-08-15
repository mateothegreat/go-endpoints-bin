// Package monitoring - provides opentelemetry tracing for the application.
package monitoring

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// Middleware is the middleware for the otel tracing service.
var Middleware echo.HandlerFunc

// Tracer instance for the otel tracing service.
var Tracer = otel.Tracer("rest.http")

// StartTracer sets up the OpenTelemetry tracing pipeline.
func StartTracer() (func(context.Context) error, error) {
	conf := config.Setup()
	ctx := context.Background()

	// Set up a trace exporter.
	traceExporter, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpoint(conf.Monitoring.Tracing.Collector), otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	hostname, _ := os.Hostname()
	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.HostName(hostname),
		semconv.ServiceName("rest"),
	))
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

// GetOtelMiddleware creates a new middleware for the echo server that creates a new span
// for each request and passes it along to the next handler.
//
// Arguments:
// - next: The next handler to call.
//
// Returns:
// - The middleware function.
func GetOtelMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	conf := config.Setup()

	if Middleware != nil {
		return Middleware
	}

	Middleware = func(c echo.Context) error {
		ctx, span := Tracer.Start(c.Request().Context(), c.Path(), trace.WithSpanKind(trace.SpanKindConsumer))
		defer span.End()

		span.SetAttributes(
			attribute.KeyValue{Key: "env", Value: attribute.StringValue(conf.Monitoring.Tracing.Collector)},
			attribute.KeyValue{Key: "containerized", Value: attribute.BoolValue(conf.Monitoring.Tracing.Enabled)},
		)

		// Pass the context with the span along to the next handler
		c.SetRequest(c.Request().WithContext(ctx))

		// Call the next handler
		return next(c)
	}

	return Middleware
}

// NewSpan creates a new span for the given context and name.
//
// Arguments:
// - ctx: The context to create the span in.
// - name: The name of the span.
// - statusCode: The status code of the span.
// - statusDescription: The status description of the span.
// - kv: The key-value pairs to set on the span.
//
// Returns:
// - The span.
func NewSpan(ctx context.Context, name string, statusCode codes.Code, statusDescription string, kv ...attribute.KeyValue) trace.Span {
	_, span := Tracer.Start(ctx, name)

	span.SetStatus(statusCode, statusDescription)
	span.SetAttributes(kv...)

	return span
}

// NewSpanWithParent creates a new span for the given context and name.
//
// Arguments:
// - ctx: The context to create the span in.
// - name: The name of the span.
//
// Returns:
// - The parent span.
// - The span.
func NewSpanWithParent(ctx context.Context, name string) (trace.Span, trace.Span) {
	parent := trace.SpanFromContext(ctx)
	_, span := Tracer.Start(ctx, name)

	return parent, span
}
