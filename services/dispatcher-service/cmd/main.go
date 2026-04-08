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
	inTopic := "threat-alerts"

	slog.Info("Dispatcher Service starting", "in_topic", inTopic)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   inTopic,
		GroupID: "dispatcher-group",
	})
	defer r.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		slog.Warn("🚨 FINAL DISPATCH: Security Incident Ready for Response", 
			"key", string(m.Key), 
			"payload", string(m.Value))
	}

	slog.Info("Dispatcher Service shutting down...")
}
