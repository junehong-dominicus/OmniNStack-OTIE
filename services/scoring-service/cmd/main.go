package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "127.0.0.1:9092"
	}
	brokers := strings.Split(brokersEnv, ",")
	inTopic := "processed-events"
	outTopic := "scored-events"

	slog.Info("Scoring Service starting", "in_topic", inTopic, "out_topic", outTopic)

	// Kafka Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   inTopic,
		GroupID: "scoring-group",
	})
	defer r.Close()

	// Kafka Writer
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    outTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		slog.Info("Enriching risk score", "offset", m.Offset, "key", string(m.Key))

		// For Phase 1 E2E, we just pass it through as a "Scored" event
		err = w.WriteMessages(ctx, kafka.Message{
			Key:   m.Key,
			Value: m.Value,
		})
		if err != nil {
			slog.Error("failed to write message", "error", err)
		}
	}

	slog.Info("Scoring Service shutting down...")
}
