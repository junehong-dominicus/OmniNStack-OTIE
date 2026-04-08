# OmniNStack Threat Intelligence Engine (OTIE) 🧠🛡️

**OmniNStack-OTIE** is a **Platform-level AI Security Engine** designed to power Zero Trust environments with Precision AI. Unlike traditional CTI (Cyber Threat Intelligence) systems, OTIE actively ingests, processes, and evaluates real-time raw telemetry (IoT, Modbus, Network, AI Cameras) using a sophisticated **Multi-Model Orchestrator**, a real-time **Feature Store**, and a **Graph-based Threat Intelligence** core.

OTIE represents the central brain of a sovereign security architecture, continuously analyzing relationships and autonomously dispatching Zero Trust defensive policies in real-time.

---

## 🏗️ 5-Layer Platform Architecture

The OTIE backend operates via a highly scalable, domain-driven microservice Monorepo representing a full ingest-to-autonomous-action pipeline:

1. **Data Fabric (`services/`, `platform/`)**: The backbone of spatial and streaming data.
   - **`collector-service`**: Highly-concurrent gRPC ingress performing localized Zero trust verification and shipping raw telemetry straight to Kafka.
   - **`processor-service`**: Normalizes and cleans streams, translating Modbus/IoT data into structured models.
   - **`feature-store`**: Streams refined data from Kafka into Redis (Online Feature Store) and Postgres (Offline Feature Store) with ultra-low latency for model consumption.
   - **`graph-db`**: Maps the connections between IP, User, Device, and Malicious payloads in real-time via Neo4j.
2. **AI Model Hub (`ai/`)**: Multi-model infrastructure (Behavior, Graph, Anomaly) collaborating through Ensemble ML.
3. **Real-time Engine**: Scores risk dynamically (`Risk = Node Risk + Edge Weight + Path Risk`).
4. **Autonomous Security (SOAR++)**: Autonomously triggers Firewall/ZTNA rule updates upon attack inference.
5. **AI Security**: Next-gen guardrails preventing GenAI prompt-injection and LLM tampering.

---

## 🚀 Local Development Environment

We provide a streamlined local environment to immediately boot your Data Fabric and test the AI-ready microservices.

### Prerequisites
- **Docker & Docker Desktop (incl. `docker-compose`)**
- **Go 1.21+**

### 1. Boot Local Data Fabric
Spin up the orchestration components (Kafka, Zookeeper, Redis, PostgreSQL, and Neo4j) natively:
```bash
docker-compose up -d
```
*Wait a few seconds for Kafka and the Graph engine to fully stabilize.*

### 2. Verify Services
Validate the architecture builds correctly across all domain packages:
```bash
go mod tidy
go build ./...
```

### 3. End-to-End Diagnostics
The repository includes synthetic data generators and native egress listeners to test the streaming backend locally:

**Terminal 1 (Listen to Kafka Egress):**
```bash
go run ./tests/integration/otie_test_conn/main.go
```

**Terminal 2 (Fire Simulated Payload):**
```bash
go run ./tests/integration/test_event_generator/main.go
```

*If synchronized properly, the synthetic Modbus payload will cleanly traverse the gRPC ingress and dump into the egress terminal with server-injected metadata.*

---

## 📚 Design Documentation

Deep-dive architecture specifications and product strategies can be found in the `design_docs/` directory:
- `01_platform_upgrade_strategy.md`: Core system roadmap and structural 5-layer capabilities.
- `02_architecture_prd.md`: Low-level service breakdown, Schema specs, and K8s Helm deployments.
- `03_pitch_deck.md`: System high-level overview and presentation script.

---

## 🛠️ Kubernetes Production Deployment
OTIE is engineered for immediate horizontal scalability. Native Helm orchestration packages reside internally (`deployments/helm`) to gracefully transition local components into highly available clusters.

```bash
helm install otie ./deployments/helm/otie
```

## 🔌 Operational CI/CD
Continuous Integration is strictly governed internally inside GitHub Actions. All workflows enforce strict unit-testing, Graph validation, and automated decoupled `Docker buildx` containerization ensuring failproof sovereign deployments.