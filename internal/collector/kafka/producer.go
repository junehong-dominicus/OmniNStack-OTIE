package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
	collector "github.com/junehong-dominicus/OmniNStack-OTIE/api/v1"
)

type Producer interface {
	PublishEvent(ctx context.Context, entry *collector.LogEntry) error
	Close() error
}

type kafkaProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(brokers []string, topic string) Producer {
	return &kafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
			Async:    true,
		},
		topic: topic,
	}
}

func (p *kafkaProducer) PublishEvent(ctx context.Context, entry *collector.LogEntry) error {
	payload, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal log entry: %w", err)
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(entry.DeviceId),
		Value: payload,
		Time:  time.Now(),
	})

	if err != nil {
		slog.Error("Failed to publish message to Kafka", "error", err, "topic", p.topic)
		return err
	}

	return nil
}

func (p *kafkaProducer) Close() error {
	return p.writer.Close()
}
