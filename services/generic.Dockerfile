# Stage 1: Build
FROM golang:1.26-alpine AS builder
RUN apk add --no-cache protoc protobuf-dev git
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
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

# Argument to specify which service to build
ARG SERVICE_PATH
RUN CGO_ENABLED=0 GOOS=linux go build -o main ${SERVICE_PATH}/cmd/main.go

# Stage 2: Final Image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
