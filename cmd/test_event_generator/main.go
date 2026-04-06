package main

import (
	"context"
	"log"
	"time"

	collector "github.com/junehong-dominicus/OmniNStack-OTIE/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("🔌 Connecting to OTIE Collector via gRPC on 127.0.0.1:50051...")
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	defer conn.Close()

	client := collector.NewCollectorServiceClient(conn)

	// Create a synthetic LogEntry payload without an ID or Timestamp
	// The CollectorService Business Logic layer should auto-generate them!
	req := &collector.IngestRequest{
		Entry: &collector.LogEntry{
			DeviceId:   "OTIE-TEST-DEVICE-01",
			MacAddress: " 00:1B:44:11:3A:B7 ", // Unsanitized MAC address
			Protocol:   "modbus",
			RawData:    `{"function_code": 3, "register": 40001, "value": [100, 200, 300]}`,
			Metadata: map[string]string{
				"environment": "local-test",
				"source":      "test_event_generator",
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("🚀 Firing Synthetic Modbus Malicious Event Payload...")
	res, err := client.Ingest(ctx, req)
	if err != nil {
		log.Fatalf("❌ gRPC Ingest Failed: %v", err)
	}

	if res.Success {
		log.Printf("✅ Payload Successfully Ingested! Server responded: %s", res.Message)
	} else {
		log.Printf("⚠️ Payload ingestion bounced! Server responded: %s", res.Message)
	}
}
