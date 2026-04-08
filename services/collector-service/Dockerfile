# Stage 1: Build & Generate
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache protoc protobuf-dev git

# Install Go protobuf plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set the working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Generate Protobuf code
RUN find api/v1 -name "*.pb.go" -delete && \
    protoc -I. --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/v1/collector.proto && \
    protoc -I. --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/v1/processor/processor.proto && \
    protoc -I. --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/v1/processor/feature.proto

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o collector services/collector-service/cmd/main.go

# Stage 2: Final Image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder Stage
COPY --from=builder /app/collector .

# Expose the gRPC port
EXPOSE 50051

# Run the binary
ENTRYPOINT ["./collector"]
