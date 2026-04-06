package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "raw-events"
	brokers := []string{"127.0.0.1:9092"}

	log.Printf("🎧 Connecting to Kafka on %v. Subscribing to topic: '%s'...", brokers, topic)

	// Create a new reader that trails the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	// Adjust offset to read the freshest messages rather than historic ones (for test visibility)
	if err := r.SetOffset(kafka.LastOffset); err != nil {
		log.Printf("⚠️ Warning: Could not set offset to latest: %v. Reading all messages.", err)
	}

	log.Println("✅ Awaiting messages! (Press Ctrl+C to quit)")

	ctx, cancel := context.WithCancel(context.Background())

	// Handle graceful shutdown
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Println("\n🛑 Stopping Kafka consumer...")
		cancel()
	}()

	for {
		// Read locally
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			log.Fatalf("❌ Error reading message: %v", err)
		}

		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Printf("📬 MSG RECEIVED AT: %v\n", time.Now().Format(time.RFC3339))
		fmt.Printf("🏷️  Key (Device ID): %s\n", string(m.Key))
		fmt.Printf("📦 Encoded Payload:\n%s\n", string(m.Value))
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("🎉 If you see UUID strings and Timestamps inside the json payload above, the Collector Business Logic successfully enriched the data!")
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
