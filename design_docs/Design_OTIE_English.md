# Design Document: OmniNStack Threat Intelligence Engine (OTIE)

## 1. High-Level Architecture Overview

The **OTIE** is built on a distributed, event-driven microservices architecture. It leverages **Kafka** as the central internal event bus to ensure decoupled communication, high throughput, and fault tolerance.

### Architecture Diagram

```mermaid
graph TD
    Client[Client/Agent/Syslog] --> AG[API Gateway]
    AG --> Auth[Auth Service]
    Auth --> AG
    AG --> Coll[Collector Service]
    Coll --> Kafka((Kafka Event Bus))
    Kafka --> Proc[Processor Service]
    Proc --> Feat[Feature Engineering]
    Feat --> ML[ML Engine]
    ML --> Score[Scoring Service]
    Score --> IOC[IOC Service]
    IOC --> Disp[Dispatcher Service]
    Disp --> External[WAF / ZTNA / SIEM]
    
    %% Storage
    Coll -.-> S3[Data Lake: S3/MinIO]
    Proc -.-> ES[Real-time: Elasticsearch]
    IOC -.-> PG[Metadata: PostgreSQL]
```

## 2. Component Design

| Service | Responsibility | Core Technology |
| :--- | :--- | :--- |
| **Auth Service** | Identity & Access (JWT, SPIFFE) | Go, Postgres |
| **Collector** | Ingestion from multiple sources | Go, Kafka Producer |
| **Processor** | Normalization & Data Cleansing | Go, Kafka Consumer/Producer |
| **Feature Service** | Behavioral feature extraction | Go/Python |
| **ML Engine** | Deep Learning Inference | Python, FastAPI, PyTorch |
| **Scoring** | Risk Score Calculation | Go |
| **IOC Service** | Repository of threat indicators | Go, Elasticsearch |
| **Dispatcher** | Automated Policy Distribution | Go, Webhooks/gRPC |

---

## 3. Sequence Diagrams (Event Flows)

### 3-1. Ingestion & Normalization Flow

```mermaid
sequenceDiagram
    participant Source as External Data Source
    participant Coll as Collector Service
    participant Kafka as Kafka (Raw Topic)
    participant Proc as Processor Service
    participant ES as Elasticsearch (Processed)

    Source->>Coll: Send Logs/Events (HTTP/Syslog)
    Coll->>Kafka: Publish Raw Event
    Kafka->>Proc: Consume Raw Event
    Proc->>Proc: Normalize & Enrich (GeoIP)
    Proc->>ES: Store Processed Log
    Proc->>Kafka: Publish Normalized Event (Processed Topic)
```

### 3-2. Threat Analysis & Scoring Flow

```mermaid
sequenceDiagram
    participant Kafka as Kafka (Processed Topic)
    participant Feat as Feature Service
    participant ML as ML Engine (GPU)
    participant Score as Scoring Service
    participant IOC as IOC Service
    participant PG as PostgreSQL

    Kafka->>Feat: Consume Processed Event
    Feat->>Feat: Compute Behavioral Features
    Feat->>ML: POST /v1/predict (Request Inference)
    ML-->>Feat: Return Prediction (Score/Anomaly)
    Feat->>Score: Submit Inference Result
    Score->>Score: Calculate Final Risk Score
    Score->>IOC: Create/Update Threat Entry
    IOC->>PG: Save Meta-intelligence
```

### 3-3. IOC Distribution & Automated Response Flow

```mermaid
sequenceDiagram
    participant IOC as IOC Service
    participant Disp as Dispatcher Service
    participant WAF as External WAF
    participant ZTNA as Zero Trust Gateway
    participant SIEM as SIEM / SOC

    IOC->>Disp: New High-Risk IOC Identified
    par Dispatch to Security Fabric
        Disp->>WAF: Block Malicious IP/URL
        Disp->>ZTNA: Revoke Session / Quaritine Device
        Disp->>SIEM: Trigger Critical Incident Alert
    end
    WAF-->>Disp: Confirmation
    ZTNA-->>Disp: Confirmation
```

---

## 4. Database Schema & ERD

The system uses **PostgreSQL** for relational metadata and **Elasticsearch** for high-volume logs/events. Below is the ERD for the PostgreSQL metadata layer.

### 4-1. PostgreSQL Entity Relationship Diagram

```mermaid
erDiagram
    USER ||--o{ AUDIT_LOG : "performs"
    THREAT_INDICATOR ||--o{ THREAT_HISTORY : "has"
    POLICY }o--|| THREAT_INDICATOR : "applies_to"
    ML_MODEL ||--o{ THREAT_INDICATOR : "identifies"
    
    USER {
        uuid id PK
        string username
        string role "Admin, Analyst, ReadOnly"
        string password_hash
        datetime last_login
    }

    THREAT_INDICATOR {
        uuid id PK
        string type "IP, URL, Domain, FileHash"
        string value
        int risk_score "0-100"
        string status "Active, Expired, FalsePositive"
        uuid model_id FK
        timestamp first_seen
        timestamp last_seen
    }

    THREAT_HISTORY {
        uuid id PK
        uuid indicator_id FK
        int old_risk_score
        int new_risk_score
        timestamp changed_at
    }

    POLICY {
        uuid id PK
        string name
        string target_system "WAF, ZTNA, SIEM"
        string action "Block, Allow, Log, Alert"
        int min_risk_to_trigger
        boolean enabled
    }

    ML_MODEL {
        uuid id PK
        string name
        string version
        string model_type "Anomaly, Classifier, Behavioral"
        float performance_score
        string artifact_path "S3_URI"
    }

    AUDIT_LOG {
        uuid id PK
        uuid user_id FK
        string action "CRUD, Login, ConfigChange"
        timestamp log_time
        string resource_id
    }
```

### 4-2. Data Storage Strategy

1. **Real-time Search (Elasticsearch)**: Used by Admin UI for log analysis and dashboard.
   - Indices: `otie-raw-events-*`, `otie-processed-events-*`.
2. **Cold Storage (S3/MinIO)**: Long-term archival for compliance and model re-training.
3. **Relational Meta-data (PostgreSQL)**: Internal system state and configuration.

---

## 5. Security & Connectivity

OTIE implements a **Zero Trust** architecture:

- **mTLS**: Every service requires mutual TLS for gRPC or HTTP communication.
- **Service Identity**: SPIFFE/SPIRE provides each pod with a verifiable SVID (Identity).
- **JWT**: Internal and External API access is controlled via short-lived JWT tokens.

---

## 6. Infrastructure & Deployment (Kubernetes)

### Helm Configuration Summary

- All deployments use **Resources Limits/Requests** for predictable performance.
- **HPA (Horizontal Pod Autoscaler)** scaling based on CPU/Memory and Kafka Lag metrics.
- **Priority Classes** used for critical services (ML Engine, Collector).

### GPU Support for ML Engine

```yaml
# Deployment Spec Fragment
resources:
  limits:
    nvidia.com/gpu: 1
```

> [!IMPORTANT]
> The ML Engine requires nodes with NVIDIA drivers and the NVIDIA Device Plugin installed for GPU acceleration.
