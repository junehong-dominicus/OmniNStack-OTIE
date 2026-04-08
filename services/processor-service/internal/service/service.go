package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
	"github.com/OmniNStack/OmniNStack-OTIE/services/processor-service/internal/kafka"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProcessorService interface {
	Start(ctx context.Context) error
	ProcessPayload(ctx context.Context, payload []byte) ([]byte, error)
}

type processorService struct {
	producer      kafka.Producer
	consumer      kafka.Consumer
	registry      map[string]ProtocolHandler
	mlEngineURL   string
	httpClient    *http.Client
}

func NewProcessorService(consumer kafka.Consumer, producer kafka.Producer) ProcessorService {
	mlURL := os.Getenv("ML_ENGINE_URL")
	if mlURL == "" {
		mlURL = "http://ml-engine-service:8000/predict"
	}

	s := &processorService{
		consumer:    consumer,
		producer:    producer,
		registry:    make(map[string]ProtocolHandler),
		mlEngineURL: mlURL,
		httpClient:  &http.Client{Timeout: 5 * time.Second},
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

	// --- INTEGRATION: Call ML Engine for Precision AI Enrichment ---
	aiInference, err := s.callMLEngine(ctx, processed)
	if err != nil {
		slog.Warn("AI Inference failed, falling back to heuristic scoring", "error", err)
	} else {
		slog.Info("AI Inference successful", "is_threat", aiInference.IsThreat, "score", aiInference.Score)
		// Update processed feature with AI results
		processed.BaselineRiskScore = aiInference.Score
		processed.RiskIndicators = append(processed.RiskIndicators, fmt.Sprintf("AI_DESC: %s", aiInference.Description))
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

type mlRequest struct {
	EventId  string    `json:"event_id"`
	Features []float64 `json:"features"`
}

type mlResponse struct {
	EventId     string  `json:"event_id"`
	Score       float32 `json:"score"`
	IsThreat    bool    `json:"is_threat"`
	Description string  `json:"description"`
}

func (s *processorService) callMLEngine(ctx context.Context, p *processor.ProcessedFeature) (*mlResponse, error) {
	// Map numerical features into a float slice for the ML model
	features := []float64{
		float64(p.BaselineRiskScore),
		float64(p.ModbusFunctionCode),
		float64(p.ModbusRegisterAddress),
	}

	reqBody, _ := json.Marshal(mlRequest{
		EventId:  p.Id,
		Features: features,
	})

	req, err := http.NewRequestWithContext(ctx, "POST", s.mlEngineURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ml-engine returned status: %d", resp.StatusCode)
	}

	var result mlResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
