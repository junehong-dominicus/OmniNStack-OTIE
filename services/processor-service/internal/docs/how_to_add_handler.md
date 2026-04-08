# How to Add a New Protocol Handler

Adding support for a new OT protocol to the OmniNStack Processor is straightforward. Follow these steps.

## Step 1: Define Proto Fields (Optional)

If the protocol has unique metadata (e.g., EtherNet/IP CIP Service) that should be promoted for ML analysis, add them to the `ProcessedFeature` message in:
`api/v1/processor/feature.proto`

**Example:**
```proto
message ProcessedFeature {
    // ... existing fields ...
    string eip_service_name = 19;
    int32 eip_attribute_id = 20;
}
```

Then regenerate the Go code (see root README for generation commands).

## Step 2: Implement the `ProtocolHandler` Interface

Create a new file in `internal/processor/service/` called `[protocol]_handler.go`. Implement the following interface:

```go
type ProtocolHandler interface {
	ProtocolName() string
	Normalize(entry *collector.LogEntry) (*processor.ProcessedFeature, error)
}
```

**Implementation Template:**

```go
package service

import (
	"encoding/json"
	"fmt"
	collector "github.com/OmniNStack/OmniNStack-OTIE/api/v1"
	processor "github.com/OmniNStack/OmniNStack-OTIE/api/v1/processor"
)

type MyData struct {
    // Define the JSON structure expected in LogEntry.RawData
}

type myHandler struct {}

func NewMyHandler() ProtocolHandler { return &myHandler{} }

func (h *myHandler) ProtocolName() string { return "my-protocol" }

func (h *myHandler) Normalize(entry *collector.LogEntry) (*processor.ProcessedFeature, error) {
    // 1. Unmarshal RawData
    // 2. Apply Zero Trust Heuristics (calculate risk score)
    // 3. Map to ProcessedFeature
    // 4. Return
}
```

## Step 3: Register the Handler

In `internal/processor/service/service.go`, add your new handler to the registry in `NewProcessorService`:

```go
func NewProcessorService(...) ProcessorService {
	s := &processorService{ ... }

	// Register Handlers
	s.registerHandler(NewModbusHandler())
	s.registerHandler(NewBACnetHandler())
	s.registerHandler(NewMyHandler()) // <--- Add your handler here

	return s
}
```

## Step 4: Verification

1.  Build the processor: `go build ./cmd/processor/main.go`
2.  Send a test event using the `test_event_generator` with the new protocol name.
3.  Verify the output in Kafka or the processor logs.
