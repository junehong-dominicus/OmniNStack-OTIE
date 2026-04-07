package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/processor/kafka"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProcessorService interface {
	Start(ctx context.Context) error
	ProcessPayload(ctx context.Context, payload []byte) ([]byte, error)
}

type processorService struct {
	producer kafka.Producer
	consumer kafka.Consumer
	registry map[string]ProtocolHandler
}

func NewProcessorService(consumer kafka.Consumer, producer kafka.Producer) ProcessorService {
	s := &processorService{
		consumer: consumer,
		producer: producer,
		registry: make(map[string]ProtocolHandler),
	}

	// Register Handlers
	s.registerHandler(NewModbusHandler())
	s.registerHandler(NewBACnetHandler())

	return s
}

func (s *processorService) registerHandler(h ProtocolHandler) {
	s.registry[h.ProtocolName()] = h
}

func (s *processorService) Start(ctx context.Context) error {
	slog.Info("Processor service starting consumption loop...")
	return s.consumer.Consume(ctx, func(ctx context.Context, key []byte, value []byte) error {
		slog.Debug("Received raw event", "key", string(key))
		
		processedPayload, err := s.ProcessPayload(ctx, value)
		if err != nil {
			slog.Error("Failed to process payload", "error", err, "key", string(key))
			return nil // Continue to next message instead of crashing
		}

		err = s.producer.PublishEvent(ctx, key, processedPayload)
		if err != nil {
			return fmt.Errorf("failed to publish processed event: %w", err)
		}

		return nil
	})
}

func (s *processorService) ProcessPayload(ctx context.Context, payload []byte) ([]byte, error) {
	// Unmarshal incoming JSON into LogEntry Protobuf
	entry := &collector.LogEntry{}
	if err := protojson.Unmarshal(payload, entry); err != nil {
		if err := json.Unmarshal(payload, entry); err != nil {
			return nil, fmt.Errorf("failed to unmarshal entry (JSON/ProtoJSON): %w", err)
		}
	}

	var processed *processor.ProcessedFeature
	var err error

	// Dispatch to handler from registry
	if handler, ok := s.registry[entry.Protocol]; ok {
		processed, err = handler.Normalize(entry)
	} else {
		// Generic fallback for unknown protocols
		processed = &processor.ProcessedFeature{
			Id:              fmt.Sprintf("proc-%s", entry.Id),
			OriginalEventId: entry.Id,
			Timestamp:       entry.Timestamp,
			DeviceId:        entry.DeviceId,
			Protocol:        entry.Protocol,
			BaselineRiskScore: 0.1,
			RiskIndicators:    []string{"UNKNOWN_PROTOCOL"},
			ProcessorVersion:  "v1.0.0-generic",
		}
	}

	if err != nil {
		return nil, fmt.Errorf("normalization failed for protocol %s: %w", entry.Protocol, err)
	}

	// Marshal processed feature to JSON for Kafka
	output, err := protojson.Marshal(processed)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal processed feature: %w", err)
	}

	slog.Info("Successfully processed event", 
		"id", entry.Id, 
		"protocol", entry.Protocol, 
		"risk_score", processed.BaselineRiskScore)
	
	return output, nil
}
