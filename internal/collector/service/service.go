package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/collector/kafka"
)

type Service interface {
	IngestLog(ctx context.Context, entry *collector.LogEntry) error
}

type collectorService struct {
	producer kafka.Producer
}

func NewCollectorService(producer kafka.Producer) Service {
	return &collectorService{
		producer: producer,
	}
}

func (s *collectorService) IngestLog(ctx context.Context, entry *collector.LogEntry) error {
	if entry == nil {
		return errors.New("log entry cannot be nil")
	}

	// 1. Validation
	if strings.TrimSpace(entry.Protocol) == "" {
		return errors.New("protocol is a mandatory field")
	}
	if strings.TrimSpace(entry.DeviceId) == "" {
		return errors.New("device_id is a mandatory field")
	}

	// 2. Data Enrichment
	if strings.TrimSpace(entry.Id) == "" {
		entry.Id = uuid.New().String()
	}
	if entry.Timestamp == nil {
		entry.Timestamp = timestamppb.Now()
	}

	// 3. Normalization
	if entry.MacAddress != "" {
		entry.MacAddress = strings.ToUpper(strings.TrimSpace(entry.MacAddress))
	}

	// 4. Produce to downstream
	return s.producer.PublishEvent(ctx, entry)
}
