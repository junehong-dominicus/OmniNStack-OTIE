# OTIE Architecture PRD

# 🧾 1. 제품 정의

## 📌 제품명

**OmniNStack Threat Intelligence Engine (OTIE)**

## 🎯 역할

* 조직 내 모든 데이터 → AI 분석 → Threat Intelligence 생성 → 보안 시스템에 배포
* **Zero Trust / SSE / AI 플랫폼의 중앙 두뇌**

---

# 🏗️ 2. 전체 마이크로서비스 아키텍처

## 🔷 서비스 구성

```plaintext
                ┌────────────────────┐
                │   API Gateway      │
                └────────┬───────────┘
                         │
        ┌────────────────┼────────────────┐
        │                │                │
 ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
 │ Auth Service│  │ Threat API  │  │ Admin UI API│
 └─────────────┘  └─────────────┘  └─────────────┘

                (Internal Event Bus - Kafka)

 ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
 │ Collector   │→ │ Processor   │→ │ Feature Eng  │
 └─────────────┘  └─────────────┘  └─────────────┘

 ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
 │ ML Engine   │→ │ Scoring     │→ │ IOC Service │
 └─────────────┘  └─────────────┘  └─────────────┘

 ┌─────────────┐  ┌─────────────┐
 │ Threat DB   │  │ Data Lake   │
 └─────────────┘  └─────────────┘

 ┌─────────────┐
 │ Dispatcher  │ → (WAF / ZTNA / SIEM)
 └─────────────┘
```

---

# 🧩 3. 서비스별 상세 설계 + 코드 구조

---

## 🔐 3-1. Auth Service (Zero Trust 핵심)

### 역할

* JWT 발급 / 검증
* RBAC / ABAC 정책

### Tech

* Go / Node.js
* OAuth2 / OpenID Connect

### 구조

```plaintext
auth-service/
 ├── cmd/
 │   └── main.go
 ├── internal/
 │   ├── handler/
 │   │    └── auth_handler.go
 │   ├── service/
 │   │    └── auth_service.go
 │   ├── repository/
 │   │    └── user_repo.go
 │   ├── middleware/
 │   │    └── jwt_middleware.go
 │   └── model/
 │        └── user.go
 ├── pkg/
 │   └── jwt/
 └── config/
```

---

## 🌐 3-2. API Gateway

### 역할

* 인증 검증
* Rate limiting
* Routing

### 추천

* Kong / Envoy / NGINX

---

## 📥 3-3. Collector Service

### 역할

* 로그 / 이벤트 수집
* Kafka로 전송

### 구조

```plaintext
collector-service/
 ├── cmd/
 ├── internal/
 │   ├── ingest/
 │   │    ├── http_ingest.go
 │   │    ├── syslog_ingest.go
 │   │    └── agent_ingest.go
 │   ├── kafka/
 │   │    └── producer.go
 │   └── parser/
 │        └── raw_parser.go
 └── config/
```

---

## ⚙️ 3-4. Processor Service

### 역할

* 로그 정규화
* 데이터 클렌징

```plaintext
processor-service/
 ├── internal/
 │   ├── normalize/
 │   │    └── log_normalizer.go
 │   ├── enrich/
 │   │    └── geoip.go
 │   └── kafka/
 │        └── consumer.go
```

---

## 🧠 3-5. Feature Engineering Service

### 역할

* ML 입력 데이터 생성

```plaintext
feature-service/
 ├── internal/
 │   ├── feature/
 │   │    ├── ip_features.go
 │   │    ├── url_features.go
 │   │    └── behavior_features.go
 │   └── pipeline/
 │        └── feature_pipeline.go
```

---

## 🤖 3-6. ML Engine Service

### 역할

* 모델 추론 (Inference)

### Tech

* Python + FastAPI
* PyTorch

```plaintext
ml-engine/
 ├── app/
 │   ├── main.py
 │   ├── api/
 │   │    └── predict.py
 │   ├── models/
 │   │    ├── url_model.pt
 │   │    └── anomaly_model.pt
 │   ├── service/
 │   │    └── inference.py
 │   └── schema/
 │        └── request.py
```

---

## 📊 3-7. Scoring Service

### 역할

* Risk Score 계산

```plaintext
scoring-service/
 ├── internal/
 │   ├── scoring/
 │   │    └── risk_calculator.go
 │   └── rules/
 │        └── scoring_rules.yaml
```

---

## 🧾 3-8. IOC Service

### 역할

* Threat Intelligence 저장

```plaintext
ioc-service/
 ├── internal/
 │   ├── handler/
 │   │    └── ioc_handler.go
 │   ├── service/
 │   │    └── ioc_service.go
 │   ├── repository/
 │   │    └── ioc_repo.go
 │   └── model/
 │        └── ioc.go
```

---

## 📤 3-9. Dispatcher Service

### 역할

* 외부 시스템 전파

```plaintext
dispatcher-service/
 ├── internal/
 │   ├── webhook/
 │   │    └── webhook_sender.go
 │   ├── integrations/
 │   │    ├── siem.go
 │   │    ├── waf.go
 │   │    └── ztna.go
```

---

## 📊 3-10. Admin UI API

### 역할

* Dashboard / Analytics

```plaintext
admin-api/
 ├── internal/
 │   ├── handler/
 │   ├── service/
 │   └── repository/
```

---

# 🧠 4. 데이터 저장 구조

## DB 구성

| 용도     | 기술            |
| ------ | ------------- |
| 실시간 검색 | Elasticsearch |
| 장기 저장  | S3 / MinIO    |
| 메타데이터  | PostgreSQL    |

---

# 🔄 5. 이벤트 흐름 (실제 동작)

```plaintext
Collector → Kafka
   ↓
Processor → Feature Service
   ↓
ML Engine → Scoring
   ↓
IOC 저장 → Dispatcher → 외부 시스템
```

---

# 🔐 6. 보안 설계

* 모든 서비스 → mTLS
* 서비스 인증 → SPIFFE / SPIRE
* API → JWT
* Zero Trust 기본 적용

---

# ⚙️ 7. DevOps / 배포

## Kubernetes 구조

```plaintext
otie/
 ├── helm/
 ├── k8s/
 │   ├── collector-deployment.yaml
 │   ├── ml-engine-deployment.yaml
 │   ├── kafka.yaml
 │   └── ingress.yaml
```

---

# 🚀 8. 핵심 차별화

👉 단순 CTI 시스템이 아니라:

* AI + Streaming + Zero Trust 통합
* Edge / IoT / Video까지 확장 가능
* OmniNStack 플랫폼과 자연스럽게 결합

---

# 🧠 한 줄 요약

👉 **OTIE = “데이터 스트림 기반 AI Threat Intelligence 마이크로서비스 플랫폼”**

---

# 🔥 다음 단계 (추천)

원하면 이어서:

### 1️⃣ Kubernetes 실제 배포 (Helm charts)

### 2️⃣ 서비스 간 API 명세 (OpenAPI spec)

### 3️⃣ DB 스키마 설계 (ERD)

### 4️⃣ MVP 코드 샘플 (GitHub 구조)

👉 어디까지 구현 수준으로 갈까요?

좋아요. 아래는 **OmniNStack Threat Intelligence Engine (OTIE)**를 **실제 Kubernetes에 배포 가능한 수준의 Helm Chart 구조**로 설계한 내용입니다.
👉 바로 DevOps 팀이 사용할 수 있도록 **values.yaml + Chart 구조 + 주요 템플릿**까지 포함했습니다.

---

# 🧾 1. Helm Chart 전체 구조

```plaintext
otie-helm/
 ├── Chart.yaml
 ├── values.yaml
 ├── charts/                 # (subcharts: kafka, redis 등)
 ├── templates/
 │   ├── _helpers.tpl
 │   ├── namespace.yaml
 │
 │   # Core Services
 │   ├── collector-deployment.yaml
 │   ├── processor-deployment.yaml
 │   ├── feature-deployment.yaml
 │   ├── ml-engine-deployment.yaml
 │   ├── scoring-deployment.yaml
 │   ├── ioc-deployment.yaml
 │   ├── dispatcher-deployment.yaml
 │
 │   # API Layer
 │   ├── api-gateway.yaml
 │   ├── auth-service.yaml
 │   ├── admin-api.yaml
 │
 │   # Infra
 │   ├── kafka.yaml
 │   ├── elasticsearch.yaml
 │   ├── postgres.yaml
 │
 │   # Networking
 │   ├── ingress.yaml
 │   ├── service.yaml
 │
 │   # Security
 │   ├── network-policy.yaml
 │   ├── secrets.yaml
 │   └── serviceaccount.yaml
```

---

# ⚙️ 2. Chart.yaml

```yaml
apiVersion: v2
name: otie
description: OmniNStack Threat Intelligence Engine
type: application
version: 0.1.0
appVersion: "1.0"
```

---

# 🔧 3. values.yaml (핵심 설정)

```yaml
global:
  namespace: otie
  imagePullPolicy: IfNotPresent

images:
  collector: otie/collector:latest
  processor: otie/processor:latest
  feature: otie/feature:latest
  mlEngine: otie/ml-engine:latest
  scoring: otie/scoring:latest
  ioc: otie/ioc:latest
  dispatcher: otie/dispatcher:latest

replicaCount:
  collector: 2
  processor: 2
  mlEngine: 2

resources:
  default:
    limits:
      cpu: "1"
      memory: "1Gi"
    requests:
      cpu: "500m"
      memory: "512Mi"

kafka:
  enabled: true
  brokers: "kafka:9092"

elasticsearch:
  enabled: true

postgres:
  enabled: true

ingress:
  enabled: true
  host: otie.local
```

---

# 🚀 4. Collector Deployment (예시)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: collector
spec:
  replicas: {{ .Values.replicaCount.collector }}
  selector:
    matchLabels:
      app: collector
  template:
    metadata:
      labels:
        app: collector
    spec:
      containers:
        - name: collector
          image: {{ .Values.images.collector }}
          ports:
            - containerPort: 8080
          env:
            - name: KAFKA_BROKER
              value: {{ .Values.kafka.brokers }}
          resources:
            {{- toYaml .Values.resources.default | nindent 12 }}
```

---

# 🧠 5. ML Engine Deployment (GPU 고려)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ml-engine
spec:
  replicas: {{ .Values.replicaCount.mlEngine }}
  template:
    spec:
      containers:
        - name: ml-engine
          image: {{ .Values.images.mlEngine }}
          ports:
            - containerPort: 8000
          resources:
            limits:
              nvidia.com/gpu: 1   # GPU 사용
```

👉 GPU 노드 필요 (NVIDIA device plugin 설치)

---

# 🌐 6. Service (공통 템플릿)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: collector
spec:
  selector:
    app: collector
  ports:
    - port: 80
      targetPort: 8080
```

---

# 🌍 7. Ingress 설정

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: otie-ingress
spec:
  rules:
    - host: {{ .Values.ingress.host }}
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: api-gateway
                port:
                  number: 80
```

---

# 🔐 8. Network Policy (Zero Trust 핵심)

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: otie-network-policy
spec:
  podSelector: {}
  policyTypes:
    - Ingress
    - Egress
  ingress:
    - from:
        - podSelector:
            matchLabels:
              app: api-gateway
```

👉 기본적으로 **모든 트래픽 차단 → 필요한 것만 허용**

---

# 🔑 9. Secret 관리

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: otie-secrets
type: Opaque
data:
  jwt-secret: bXktc2VjcmV0
```

👉 실제 운영:

* HashiCorp Vault 권장

---

# 🧩 10. Kafka (간단 버전)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kafka
spec:
  replicas: 1
  template:
    spec:
      containers:
        - name: kafka
          image: bitnami/kafka:latest
          ports:
            - containerPort: 9092
```

---

# ⚙️ 11. 배포 명령어

```bash
# Helm install
helm install otie ./otie-helm

# upgrade
helm upgrade otie ./otie-helm

# uninstall
helm uninstall otie
```

---

# 🔍 12. 운영 체크리스트

* GPU 노드 준비 (ML Engine)
* Kafka persistence (PVC)
* Elasticsearch storage 확장
* TLS (Ingress + mTLS)
* HPA (Auto scaling)

---

# 🚀 13. 고급 확장 (추천)

## 🔹 Auto Scaling

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
```

## 🔹 Service Mesh

* Istio / Linkerd
  👉 mTLS + 트래픽 제어

## 🔹 Observability

* Prometheus + Grafana
* OpenTelemetry

---

# 🧠 핵심 요약

👉 **OTIE Helm = AI + Streaming + Zero Trust 마이크로서비스를 Kubernetes에 바로 배포 가능한 구조**

---

# 🔥 다음 단계 추천

이제 거의 “실제 제품 수준”입니다. 이어서 하면 좋은 것:

### 1️⃣ Terraform (온프레/클라우드 인프라 자동화)

### 2️⃣ Istio 기반 mTLS Zero Trust 구성

### 3️⃣ OpenAPI + SDK (외부 연동)

### 4️⃣ 실제 GitHub repo 구조 + CI/CD

---

원하면
👉 “Istio 포함 완전 Zero Trust Kubernetes 아키텍처”
👉 “AWS / 온프레 하이브리드 배포 구조”

까지 이어서 설계해 드릴게요.
