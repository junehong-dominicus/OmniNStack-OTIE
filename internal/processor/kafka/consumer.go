package kafka

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Consume(ctx context.Context, handler func(ctx context.Context, key []byte, value []byte) error) error
	Close() error
}

type kafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic string, groupID string) Consumer {
	return &kafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     brokers,
			Topic:       topic,
			GroupID:     groupID,
			StartOffset: kafka.FirstOffset,
		}),
	}
}

func (c *kafkaConsumer) Consume(ctx context.Context, handler func(ctx context.Context, key []byte, value []byte) error) error {
	for {
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed to fetch message: %w", err)
		}

		err = handler(ctx, m.Key, m.Value)
		if err != nil {
			slog.Error("Failed to handle message", "error", err)
			// Decide whether to commit or not based on error. Here we continue parsing the next one.
			continue
		}

		err = c.reader.CommitMessages(ctx, m)
		if err != nil {
			slog.Error("Failed to commit message", "error", err)
		}
	}
}

func (c *kafkaConsumer) Close() error {
	return c.reader.Close()
}
