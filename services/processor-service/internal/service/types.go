package service

import (
	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
)

// ProtocolHandler is the interface that all OT protocol normalization sub-services must implement.
type ProtocolHandler interface {
	// ProtocolName returns the name of the protocol this handler can process (e.g., "modbus", "bacnet").
	ProtocolName() string

	// Normalize processes the raw LogEntry into a structured ProcessedFeature with risk indicators.
	Normalize(entry *collector.LogEntry) (*processor.ProcessedFeature, error)
}
