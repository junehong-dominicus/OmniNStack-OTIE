package handler

import (
	"context"
	"log/slog"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	"github.com/OmniNStack/OmniNStack-OTIE/services/collector-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectorHandler struct {
	collector.UnimplementedCollectorServiceServer
	svc service.Service
}

func NewCollectorHandler(svc service.Service) *CollectorHandler {
	return &CollectorHandler{
		svc: svc,
	}
}

func (h *CollectorHandler) Ingest(ctx context.Context, req *collector.IngestRequest) (*collector.IngestResponse, error) {
	slog.Info("Ingesting log entry", "device_id", req.Entry.DeviceId, "protocol", req.Entry.Protocol)

	err := h.svc.IngestLog(ctx, req.Entry)
	if err != nil {
		slog.Error("Failed to ingest log entry", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "Ingestion failed: %v", err)
	}

	return &collector.IngestResponse{
		Success: true,
		Message: "Log entry successfully ingested",
	}, nil
}
