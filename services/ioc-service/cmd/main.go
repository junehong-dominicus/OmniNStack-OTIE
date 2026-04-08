package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/segmentio/kafka-go"
)

type ProcessedFeature struct {
	Id                string  `json:"id"`
	BaselineRiskScore float32 `json:"baselineRiskScore"`
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "127.0.0.1:9092"
	}
	brokers := strings.Split(brokersEnv, ",")
	inTopic := "scored-events"
	outTopic := "threat-alerts"

	slog.Info("IOC Service starting", "in_topic", inTopic, "out_topic", outTopic)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   inTopic,
		GroupID: "ioc-group",
	})
	defer r.Close()

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

		var feature ProcessedFeature
		if err := json.Unmarshal(m.Value, &feature); err != nil {
			slog.Warn("Failed to parse event", "error", err)
			continue
		}

		// Heuristic: If risk score > 0.5, it's a threat alert
		if feature.BaselineRiskScore > 0.5 {
			slog.Warn("⚠️ THREAT DETECTED", "event_id", feature.Id, "score", feature.BaselineRiskScore)
			err = w.WriteMessages(ctx, kafka.Message{
				Key:   m.Key,
				Value: m.Value,
			})
			if err != nil {
				slog.Error("failed to write alert", "error", err)
			}
		} else {
			slog.Info("Event verified clean", "event_id", feature.Id, "score", feature.BaselineRiskScore)
		}
	}

	slog.Info("IOC Service shutting down...")
}
