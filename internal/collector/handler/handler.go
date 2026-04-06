package handler

import (
	"context"
	"log/slog"

	"github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/collector/kafka"
)

type CollectorHandler struct {
	collector.UnimplementedCollectorServiceServer
	producer kafka.Producer
}

func NewCollectorHandler(producer kafka.Producer) *CollectorHandler {
	return &CollectorHandler{
		producer: producer,
	}
}

func (h *CollectorHandler) Ingest(ctx context.Context, req *collector.IngestRequest) (*collector.IngestResponse, error) {
	slog.Info("Ingesting log entry", "device_id", req.Entry.DeviceId, "protocol", req.Entry.Protocol)

	err := h.producer.PublishEvent(ctx, req.Entry)
	if err != nil {
		slog.Error("Failed to ingest log entry", "error", err)
		return &collector.IngestResponse{
			Success: false,
			Message: "Failed to publish to Kafka",
		}, nil
	}

	return &collector.IngestResponse{
		Success: true,
		Message: "Log entry successfully ingested",
	}, nil
}
