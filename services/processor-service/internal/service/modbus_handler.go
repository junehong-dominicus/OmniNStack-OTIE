package service

import (
	"encoding/json"
	"fmt"

	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
)

// ModbusData represents the expected JSON structure in LogEntry.RawData for Modbus
type ModbusData struct {
	FunctionCode int32   `json:"function_code"`
	Register     int32   `json:"register"`
	Value        []int32 `json:"value"`
}

type modbusHandler struct {
	version string
}

func NewModbusHandler() ProtocolHandler {
	return &modbusHandler{
		version: "v1.0.0-zt",
	}
}

func (h *modbusHandler) ProtocolName() string {
	return "modbus"
}

func (h *modbusHandler) Normalize(entry *collector.LogEntry) (*processor.ProcessedFeature, error) {
	var data ModbusData
	if err := json.Unmarshal([]byte(entry.RawData), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal modbus data: %w", err)
	}

	riskIndicators := []string{}
	baselineRisk := float32(0.0)

	// --- Zero Trust Heuristic Rules ---

	// Rule MODBUS-01: Function Code Validation (Least Privilege)
	allowedCodes := map[int32]bool{3: true, 4: true}
	if !allowedCodes[data.FunctionCode] {
		riskIndicators = append(riskIndicators, "ZT_MODBUS_UNUSUAL_FUNC")
		baselineRisk += 0.4
	}

	// Rule MODBUS-02: Register Range Profiling
	if data.Register > 40100 {
		riskIndicators = append(riskIndicators, "ZT_MODBUS_OUT_OF_BOUNDS")
		baselineRisk += 0.3
	}

	// Rule MODBUS-03: Value Anomaly
	for _, v := range data.Value {
		if v > 1000 {
			riskIndicators = append(riskIndicators, "ZT_MODBUS_VALUE_SPIKE")
			baselineRisk += 0.2
			break
		}
	}

	if baselineRisk > 1.0 {
		baselineRisk = 1.0
	}

	return &processor.ProcessedFeature{
		Id:                    fmt.Sprintf("proc-%s", entry.Id),
		OriginalEventId:       entry.Id,
		Timestamp:             entry.Timestamp,
		DeviceId:              entry.DeviceId,
		Protocol:              h.ProtocolName(),
		ModbusFunctionCode:    data.FunctionCode,
		ModbusRegisterAddress: data.Register,
		ModbusValues:          data.Value,
		BaselineRiskScore:     baselineRisk,
		RiskIndicators:        riskIndicators,
		ProcessorVersion:      h.version,
	}, nil
}
