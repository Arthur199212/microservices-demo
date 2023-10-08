package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	"github.com/Arthur199212/microservices-demo/src/shipping/gapi"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
)

const (
	defaultPort = "5004"
	serviceName = "shipping"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if os.Getenv("ENABLE_TRACING") == "1" {
		log.Info().Msg("tracing is enabled")
		tp := initTracerProvider()
		defer func() {
			if err := tp.Shutdown(context.Background()); err != nil {
				log.Err(err).Msg("error shutting down tracer provider")
			}
		}()
	} else {
		log.Info().Msg("tracing is disabled")
	}
	tracer := otel.Tracer(serviceName)

	shippingService := shipping.NewShippingService()
	srv := gapi.NewServer(tracer, shippingService)
	grpcServer := grpc.NewServer()
	shippingv1.RegisterShippingServiceServer(grpcServer, srv)
	grpc_health_v1.RegisterHealthServer(grpcServer, srv)

	// to provide self-documentation
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	go func() {
		log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot start gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("server is shutting down ...")

	grpcServer.GracefulStop()
}

func initTracerProvider() *sdktrace.TracerProvider {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var collectorAddr string
	mustMapEnv(&collectorAddr, "COLLECTOR_SERVICE_ADDR")
	var collectorConn *grpc.ClientConn
	mustConnGRPC(ctx, &collectorConn, collectorAddr)

	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithGRPCConn(collectorConn),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create new otlp trace grpc exporter")
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(initResource()),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)
	// Propagate trace context
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func initResource() *sdkresource.Resource {
	extraResources, err := sdkresource.New(
		context.Background(),
		sdkresource.WithAttributes(
			semconv.ServiceName(serviceName),
			attribute.String("environment", "demo"),
		),
		sdkresource.WithSchemaURL(semconv.SchemaURL),
		sdkresource.WithProcess(),
		sdkresource.WithOS(),
		sdkresource.WithContainer(),
		sdkresource.WithHost(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create extra resources")
	}
	resource, err := sdkresource.Merge(
		sdkresource.Default(),
		extraResources,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create new otlp resource")
	}
	return resource
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Fatal().Msg(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	var err error
	*conn, err = grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create gRPC connection to %q", addr)
	}
}
