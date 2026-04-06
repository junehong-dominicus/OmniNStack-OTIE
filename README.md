# OmniNStack Threat Intelligence Engine (OTIE) 🛡️

**OmniNStack-OTIE** is advanced machine-learning driven threat intelligence architecture designed to actively ingest, process, enrich, and evaluate raw telemetry from IoT devices (Modbus, Syslog, Webhooks). Written primarily in Go, OTIE deploys a highly scalable microservice lattice utilizing gRPC and Apache Kafka to stream high-velocity data down an ML-powered inference pipeline.

## 🏗️ Core Architecture
The OTIE backend operates via a decoupled, sovereign microservice architecture representing a full ingest-to-IOC data pipeline:

1. **`collector-service`**: *(Currently Active)* - The front-line ingress point. Receives raw device telemetry payload streams via highly efficient `gRPC` endpoints, performs zero-trust ID validation, automatically enriches data loops with server-side generation mechanisms (UUID formulation, Timestamp injection), and securely streams payloads directly onto the `raw-events` Apache Kafka backbone.
2. **`processor-service`**: Consumes from Kafka, sanitizes data structures, and normalizes payloads into ML-parseable feature logic.
3. **`ml_engine-service`**: Analyzes the generated event vectors using proprietary threat detection telemetry models.
4. **`scoring-service`**: Dynamically evaluates the contextualized ML inference results and attributes explicit cyber-threat Risk Scores.
5. **`ioc_service-service`**: Persistently tracks Indicators of Compromise inside a localized PostgreSQL enclave and orchestrates dynamic dispatcher webhook alerts.

---

## 🚀 Local Development Environment

### Prerequisites
- **Git**
- **Docker & Docker Desktop (incl. `docker-compose`)**
- **Go 1.26+**

### Infrastructure Deployment
Boot up the core orchestration components (including Apache Kafka, Elasticsearch, PostgreSQL, and the internal OTIE backend Collector container) leveraging Docker:
```bash
docker-compose up -d --build
```
*Wait approximately 5-10 seconds for Kafka KRaft clusters to natively stabilize.*

### E2E Diagnostic Tools
The project ships with native Go diagnostic tooling meant to aggressively circumvent integration bottlenecks. To safely test the Collector pipeline locally:

1. **Spin up the Network Listener** (Egress validation tool binding natively to our Kafka ingestion topic):
```bash
go run ./cmd/otie_test_conn/main.go
```
2. **Fire the Simulated Payload** (gRPC synthetic client firing a mocked malicious Modbus event):
```bash
go run ./cmd/test_event_generator/main.go
```
*If everything is synchronized smoothly, you will observe the mocked payload dump beautifully into the Network listener tab, cleanly appended with Collector-generated UUIDs and Server timestamps.*

---

## 🛠️ Kubernetes Production Deployment
OTIE is engineered for immediate horizontal scalability. Native Helm orchestration packages reside internally to gracefully transition your development clusters into active Kubernetes deployments.

```bash
helm install collector ./deployments/helm/collector \
  --set image.repository=github.com/junehong-dominicus/OmniNStack-OTIE
```

---

## 🔌 Operational CI/CD
Continuous Integration is strictly governed natively inside GitHub Actions (`.github/workflows/ci.yaml`). All active workflow matrices invoke isolated `golang:1.26` runners performing strict unit testing bounds alongside virtual asynchronous Docker `Buildx` verifications ensuring the core logic remains completely failproof. User Docker Registry destination overrides can be actively configured via dynamic UI `workflow_dispatch` mechanisms.