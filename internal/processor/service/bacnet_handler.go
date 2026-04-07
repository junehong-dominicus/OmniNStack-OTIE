package service

import (
	"encoding/json"
	"fmt"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
)

// BACnetData represents the expected JSON structure in LogEntry.RawData for BACnet
type BACnetData struct {
	ServiceChoice int32  `json:"service_choice"`
	ObjectType    string `json:"object_type"`
	Instance      int32  `json:"instance"`
	Property      string `json:"property"`
	Priority      int32  `json:"priority"`
	Value         string `json:"value"`
}

type bacnetHandler struct {
	version string
}

func NewBACnetHandler() ProtocolHandler {
	return &bacnetHandler{
		version: "v1.0.0-zt",
	}
}

func (h *bacnetHandler) ProtocolName() string {
	return "bacnet"
}

func (h *bacnetHandler) Normalize(entry *collector.LogEntry) (*processor.ProcessedFeature, error) {
	var data BACnetData
	if err := json.Unmarshal([]byte(entry.RawData), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal bacnet data: %w", err)
	}

	riskIndicators := []string{}
	baselineRisk := float32(0.0)

	// --- Zero Trust Heuristic Rules ---

	// Rule BACNET-01: Service Choice Restriction (Least Privilege)
	// Service 15 (WriteProperty), 16 (WritePropertyMultiple) are restricted.
	if data.ServiceChoice == 15 || data.ServiceChoice == 16 {
		riskIndicators = append(riskIndicators, "ZT_BACNET_UNAUTHORIZED_WRITE")
		baselineRisk += 0.4
	}

	// Rule BACNET-02: Safety Priority Override
	// Priority 1 (Manual-Life Safety), 2 (Automatic-Life Safety), 3 (Available)
	if data.Priority > 0 && data.Priority <= 3 {
		riskIndicators = append(riskIndicators, "ZT_BACNET_SAFETY_PRIORITY_OVERRIDE")
		baselineRisk += 0.5
	}

	// Rule BACNET-03: Object Instance Out-of-Bounds
	// Example: Typical deployment uses instances 0-1000 for sensors.
	if data.Instance > 1000 {
		riskIndicators = append(riskIndicators, "ZT_BACNET_INSTANCE_OUT_OF_RANGE")
		baselineRisk += 0.2
	}

	// Rule BACNET-04: Disruption Service Detection
	// Service 20 (ReinitializeDevice), 6 (AtomicWriteFile)
	if data.ServiceChoice == 20 || data.ServiceChoice == 6 {
		riskIndicators = append(riskIndicators, "ZT_BACNET_DISRUPTION_SERVICE")
		baselineRisk += 0.8
	}

	if baselineRisk > 1.0 {
		baselineRisk = 1.0
	}

	return &processor.ProcessedFeature{
		Id:                       fmt.Sprintf("proc-%s", entry.Id),
		OriginalEventId:          entry.Id,
		Timestamp:                entry.Timestamp,
		DeviceId:                 entry.DeviceId,
		Protocol:                 h.ProtocolName(),
		BacnetObjectType:         data.ObjectType,
		BacnetObjectInstance:     data.Instance,
		BacnetPropertyIdentifier: data.Property,
		BacnetPriority:           data.Priority,
		BaselineRiskScore:        baselineRisk,
		RiskIndicators:           riskIndicators,
		ProcessorVersion:         h.version,
	}, nil
}
