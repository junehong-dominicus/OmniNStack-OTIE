package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const registryURL = "http://localhost:8081"

type SchemaPayload struct {
	Schema     string `json:"schema"`
	SchemaType string `json:"schemaType"`
}

type RegistrationResponse struct {
	ID      int `json:"id"`
	Version int `json:"version"`
}

func registerSchema(subject, filePath string) error {
	fmt.Printf("📡 Registering schema for subject: %s from %s...\n", subject, filePath)

	// Read the proto file content
	schemaContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	payload := SchemaPayload{
		Schema:     string(schemaContent),
		SchemaType: "PROTOBUF",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/subjects/%s/versions", registryURL, subject)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/vnd.schemaregistry.v1+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("registration failed (%d): %s", resp.StatusCode, string(body))
	}

	var regResp RegistrationResponse
	if err := json.Unmarshal(body, &regResp); err != nil {
		// Version might be missing in some responses, fall back to ID
		fmt.Printf("✅ Successfully registered! Response: %s\n", string(body))
	} else {
		fmt.Printf("✅ Successfully registered! Version: %d (ID: %d)\n", regResp.Version, regResp.ID)
	}

	return nil
}

func main() {
	schemas := []struct {
		subject string
		path    string
	}{
		{"raw-events-value", "api/v1/collector.proto"},
		{"processed-events-value", "api/v1/processor/feature.proto"},
	}

	for _, s := range schemas {
		if err := registerSchema(s.subject, s.path); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error registering %s: %v\n", s.subject, err)
			os.Exit(1)
		}
	}

	fmt.Println("🎉 All schemas registered!")
}
