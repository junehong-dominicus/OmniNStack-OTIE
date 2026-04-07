package kafka

import (
	"context"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishEvent(ctx context.Context, key []byte, payload []byte) error
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
			AllowAutoTopicCreation: true,
		},
		topic: topic,
	}
}

func (p *kafkaProducer) PublishEvent(ctx context.Context, key []byte, payload []byte) error {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
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
