package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/processor/handler"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/processor/kafka"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/processor/service"

	"google.golang.org/grpc"
)

func main() {
	// Initialize logging
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Kafka Configuration
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "127.0.0.1:9092"
	}
	kafkaBrokers := strings.Split(brokersEnv, ",")

	inTopic := os.Getenv("KAFKA_IN_TOPIC")
	if inTopic == "" {
		inTopic = "raw-events"
	}

	outTopic := os.Getenv("KAFKA_OUT_TOPIC")
	if outTopic == "" {
		outTopic = "processed-events"
	}

	groupID := "processor-service-group"

	// Initialize Kafka Producer and Consumer
	producer := kafka.NewKafkaProducer(kafkaBrokers, outTopic)
	defer producer.Close()

	consumer := kafka.NewKafkaConsumer(kafkaBrokers, inTopic, groupID)
	defer consumer.Close()

	// Initialize Business Logic Service Layer
	processorService := service.NewProcessorService(consumer, producer)

	// Start Background Process Loop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := processorService.Start(ctx); err != nil {
			slog.Error("Processor background loop exited with error", "error", err)
		}
	}()

	// Initialize gRPC Handler
	processorHandler := handler.NewProcessorHandler(processorService)

	// Start gRPC Server
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50052"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	processor.RegisterProcessorServiceServer(grpcServer, processorHandler)

	slog.Info("Processor Service starting on :" + port)

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

	slog.Info("Processor Service shutting down...")
	cancel() // Cancel background consumer context
	grpcServer.GracefulStop()
}
