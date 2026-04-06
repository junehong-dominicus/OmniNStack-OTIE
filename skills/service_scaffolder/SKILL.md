# Skill: OTIE Go Service Scaffolder

## Overview

This skill automates the creation of a new Go-based microservice within the OTIE platform, following the structure and guidelines defined in the Design documentation.

## Guidelines

-   **Project Structure**:
    -   `cmd/{service}/main.go`: Entry point for binary.
    -   `internal/{service}/handler/`: gRPC/HTTP handlers.
    -   `internal/{service}/service/`: Business logic.
    -   `internal/{service}/repository/`: Database/Storage interactions.
    -   `internal/{service}/model/`: Domain entities.
    -   `pkg/`: (optional) Shared utility libraries for the service.

-   **Dependencies**:
    -   `google.golang.org/grpc` for communication.
    -   `github.com/segmentio/kafka-go` for event bus interaction.
    -   `github.com/jackc/pgx/v5` for PostgreSQL interaction.
    -   `log/slog` for structured logging.

-   **Initialization Flow**:
    1.  Create the `cmd` and `internal` directories.
    2.  Copy the Protobuf client/server libraries from `pkg/generated/`.
    3.  Initialize a `go.mod` file with necessary dependencies.
    4.  Scaffold a basic `main.go` with Kafka producer/consumer setup.
    5.  Generate a `Dockerfile` for containerization.
    6.  Generate a basic Helm chart in `deployments/helm/`.

## Scripting Template

```bash
# Example command (scaffold a new service)
OTIE_ScaffoldService collector
```
