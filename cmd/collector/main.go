package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	collector "github.com/junehong-dominicus/OmniNStack-OTIE/api/v1"
	"github.com/junehong-dominicus/OmniNStack-OTIE/internal/collector/handler"
	"github.com/junehong-dominicus/OmniNStack-OTIE/internal/collector/kafka"
	"github.com/junehong-dominicus/OmniNStack-OTIE/internal/collector/service"

	"google.golang.org/grpc"
)

func main() {
	// Initialize logging
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Initialize Kafka Producer
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "127.0.0.1:9092"
	}
	kafkaBrokers := strings.Split(brokersEnv, ",")
	kafkaTopic := "raw-events"
	producer := kafka.NewKafkaProducer(kafkaBrokers, kafkaTopic)
	defer producer.Close()

	// Initialize Business Logic Service Layer
	collectorService := service.NewCollectorService(producer)

	// Initialize gRPC Handler
	collectorHandler := handler.NewCollectorHandler(collectorService)

	// Start gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	collector.RegisterCollectorServiceServer(grpcServer, collectorHandler)

	slog.Info("Collector Service starting on :50051")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to serve", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("Collector Service shutting down...")
	grpcServer.GracefulStop()
}
