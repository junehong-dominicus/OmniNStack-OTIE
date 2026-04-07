package handler

import (
	"context"

	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
	"github.com/OmniNStack/OmniNStack-OTIE/internal/processor/service"
)

type ProcessorHandler struct {
	processor.UnimplementedProcessorServiceServer
	svc service.ProcessorService
}

func NewProcessorHandler(svc service.ProcessorService) *ProcessorHandler {
	return &ProcessorHandler{
		svc: svc,
	}
}

func (h *ProcessorHandler) Process(ctx context.Context, req *processor.ProcessRequest) (*processor.ProcessResponse, error) {
	// Exposes the core processing logic to a synchronous gRPC method
	processed, err := h.svc.ProcessPayload(ctx, []byte(req.RawData))
	if err != nil {
		return &processor.ProcessResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &processor.ProcessResponse{
		Success:        true,
		Message:        "Successfully processed payload",
		ProcessedData:  string(processed),
	}, nil
}

func (h *ProcessorHandler) CheckHealth(ctx context.Context, req *processor.HealthCheckRequest) (*processor.HealthCheckResponse, error) {
	return &processor.HealthCheckResponse{
		IsHealthy: true,
		Status:    "Processor Service is running and accepting events",
	}, nil
}
