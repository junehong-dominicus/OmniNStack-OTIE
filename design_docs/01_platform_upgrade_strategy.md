# OTIE Platform Upgrade Strategy

# 🧠 0. 목표 정의 (업그레이드 방향)

기존 OTIE
→ “CTI 엔진”

업그레이드 OTIE
→ **“Platform-level AI Security Engine”**

참고 개념:

* Palo Alto Networks Precision AI
* Cortex XSIAM

---

# 🏗️ 1. 전체 아키텍처 (Precision AI급)

```plaintext
                ┌──────────────────────────────┐
                │     OmniNStack Platform      │
                │ (VMS / AI / Zero Trust / IoT)│
                └─────────────┬────────────────┘
                              ↓
                ┌──────────────────────────────┐
                │        OTIE AI Core          │
                │  (Multi-Model Intelligence)  │
                └─────────────┬────────────────┘

 ┌───────────────┬────────────┼───────────────┬──────────────┐
 ↓               ↓            ↓               ↓              ↓

[Data Fabric] [Model Hub] [Real-time Engine] [Automation] [AI Security]

```

👉 핵심:
**OTIE = 중앙 AI 엔진 + 모든 시스템 연결**

---

# 🧩 2. 핵심 모듈 (5 Layer 구조)

---

## 🧱 2-1. Data Fabric Layer (가장 중요)

👉 Precision AI의 핵심 = “데이터”

### 구성

```plaintext
[Data Sources]
- Network
- Endpoint
- AI Camera / IoT
- Cloud / SaaS
- Threat Feeds

        ↓

[Unified Data Fabric]
- Stream (Kafka)
- Data Lake (S3/MinIO)
- Feature Store
- Graph DB (관계 분석)
```

### 핵심 기능

* 실시간 + 배치 통합
* 데이터 정규화
* 엔티티 그래프 (IP, User, Device 연결)

👉 업그레이드 포인트:

* 단순 로그 → **관계 기반 데이터**

---

## 🤖 2-2. Model Hub (Multi-Model AI)

👉 Precision AI 핵심 = “모델 다양성”

### 구성

```plaintext
[Model Registry]
 - URL Detection
 - Malware
 - Behavior
 - Graph AI
 - GenAI

[Model Orchestrator]
 - 모델 선택
 - Ensemble 처리
```

### 특징

* 단일 모델 ❌
* Multi-model ensemble ✅

### 추가 (중요)

* Online learning
* Model versioning
* A/B testing

---

## ⚡ 2-3. Real-time Intelligence Engine

👉 실시간 분석 엔진

```plaintext
Event Stream → Feature → Model → Scoring → Decision
```

### 기능

* sub-second detection
* correlation (이벤트 연결)
* anomaly detection

### 기술

* Flink / Spark Streaming

---

## 🤖 2-4. Autonomous Security (SOAR++)

👉 Precision AI 핵심 차별점

### 기존

* Alert → 사람이 대응

### 업그레이드

* Alert → AI가 대응

---

### 구성

```plaintext
[Decision Engine]
   ↓
[Playbook Engine]
   ↓
[Action Engine]
```

### 자동 대응 예시

* ZTNA 접근 차단
* WAF rule 생성
* Endpoint 격리

---

## 🧬 2-5. AI Security Layer (GenAI 보호)

👉 최신 트렌드 (중요)

### 보호 대상

* AI 모델
* AI API
* LLM Agent

### 기능

* Prompt Injection 탐지
* Data leakage 방지
* AI 사용 가시성

---

# 🔄 3. 데이터 → AI → 자동화 흐름

```plaintext
[All Data Sources]
        ↓
[Data Fabric]
        ↓
[Multi-model AI]
        ↓
[Risk Scoring + Correlation]
        ↓
[Autonomous Response]
        ↓
[System Enforcement]
```

👉 완전 자동 루프

---

# 🧠 4. 핵심 기술 업그레이드

## 기존 OTIE vs 업그레이드

| 영역  | 기존          | 업그레이드                     |
| --- | ----------- | ------------------------- |
| 데이터 | 로그 중심       | Graph + Feature Store     |
| AI  | 단일 모델       | Multi-model orchestration |
| 처리  | 배치 + 일부 실시간 | Full real-time            |
| 대응  | 수동          | Autonomous                |
| 범위  | CTI         | Platform-wide             |

---

# 🔐 5. Zero Trust 통합

OTIE는 단독이 아니라:

```plaintext
OTIE → Policy Engine → ZTNA / IAM / WAF
```

👉 모든 보안 정책은 **OTIE 기반으로 동적 생성**

---

# 📊 6. SOC 자동화 (Cortex 스타일)

### 기능

* Incident 자동 생성
* Root cause 분석
* 자동 remediation

### GenAI 활용

* 자연어로 분석
* “왜 공격인지 설명”

---

# 🧠 7. 핵심 알고리즘 (중요)

## 1. Graph-based Threat Detection

* IP ↔ User ↔ Device 관계 분석

## 2. Ensemble ML

* 여러 모델 결과 결합

## 3. Behavior Profiling

* 정상 vs 비정상 패턴

## 4. Risk Scoring

```plaintext
Risk = (Threat Score × Confidence) + Context
```

---

# 🚀 8. OmniNStack 통합 구조

```plaintext
[AI Camera]
[IoT]
[VMS]
[User Activity]

        ↓

      OTIE

        ↓

[Zero Trust Enforcement]
[Security Automation]
```

👉 AI + 보안 완전 통합

---

# 🔥 9. 차별화 전략 (사업 관점)

## Precision AI vs OTIE

| 항목  | Precision AI | OTIE             |
| --- | ------------ | ---------------- |
| 데이터 | 글로벌          | 고객 + Edge        |
| 강점  | 규모           | 온프레 + AI 영상      |
| 차별화 | 플랫폼          | Edge + AI Camera |

👉 OTIE 포지션:
**“Edge + AI + Zero Trust 통합 AI 보안 플랫폼”**

---

# 🧠 10. 한 줄 정의

👉 **OTIE = “Data Fabric + Multi-AI + Autonomous Security” 기반 플랫폼 보안 엔진**

---

# 🔥 다음 단계 (강력 추천)

이제 거의 “엔터프라이즈 제품 수준”입니다.

다음으로 가면 완성됩니다:

### 1️⃣ 실제 서비스별 상세 API 설계 (OpenAPI)

### 2️⃣ Graph DB 스키마 설계 (핵심)

### 3️⃣ AI 모델 설계 (학습 데이터 + 구조)

### 4️⃣ 투자용 Pitch Deck

---

👉 추천 다음 단계:
**“Graph 기반 Threat Intelligence DB 설계” (Precision AI 핵심)**

좋아요. 이제 **OTIE를 Precision AI급으로 만드는 핵심 요소**인
👉 **Graph 기반 Threat Intelligence DB 설계**를 *실제 구현 가능한 수준*으로 정리합니다.

---

# 🧠 1. 왜 Graph DB인가 (핵심 이유)

기존 (Relational / Log 중심) ❌
→ “이벤트 단위 분석”

Graph 기반 ✅
→ **“관계(Relationship) 기반 공격 탐지”**

예:

* 같은 IP → 여러 계정 로그인
* 특정 파일 → 여러 디바이스 감염
* 사용자 → 의심 도메인 반복 접근

👉 이런 **연결성 공격 패턴**은 Graph가 아니면 탐지 어려움

---

# 🏗️ 2. 전체 구조 (OTIE Graph Layer)

```plaintext id="p1ykkk"
[Raw Events]
   ↓
[Entity Extraction]
   ↓
[Graph Builder]
   ↓
[Graph DB]
   ↓
[Graph AI / Query Engine]
```

---

# 🧩 3. Graph 데이터 모델 (핵심 설계)

## 🔷 Node (엔티티)

| Node Type | 설명                |
| --------- | ----------------- |
| User      | 사용자               |
| Device    | PC / Mobile / IoT |
| IP        | IP 주소             |
| Domain    | 도메인               |
| URL       | URL               |
| File      | 파일 (hash)         |
| Process   | 실행 프로세스           |
| AI Model  | LLM / API         |
| Camera    | AI Camera         |

---

## 🔗 Edge (관계)

| Edge           | 의미                |
| -------------- | ----------------- |
| LOGGED_IN_FROM | User → IP         |
| ACCESSED       | User → URL        |
| CONNECTED_TO   | Device → IP       |
| DOWNLOADED     | Device → File     |
| EXECUTED       | Device → Process  |
| RESOLVED_TO    | Domain → IP       |
| GENERATED_BY   | AI Model → Output |
| STREAMS_TO     | Camera → Server   |

---

## 📌 예시 Graph

```plaintext id="pnq20z"
(User A)
   │ LOGGED_IN_FROM
   ↓
 (IP 1.2.3.4)
   │ CONNECTED_TO
   ↓
 (Malicious Domain)
   │ HOSTS
   ↓
 (Malware File)
```

👉 공격 체인 완성

---

# 🧠 4. 핵심 스키마 (Neo4j 기준)

## Node 정의

```cypher id="gs5q0n"
CREATE (u:User {id: "user123"})
CREATE (ip:IP {address: "1.2.3.4"})
CREATE (d:Domain {name: "bad.com"})
```

---

## Relationship 생성

```cypher id="vybtw6"
MATCH (u:User {id: "user123"}), (ip:IP {address: "1.2.3.4"})
CREATE (u)-[:LOGGED_IN_FROM]->(ip)
```

---

# ⚡ 5. Graph Query (탐지 로직 핵심)

## 🎯 1. Suspicious Login Pattern

```cypher id="w9r9ci"
MATCH (u:User)-[:LOGGED_IN_FROM]->(ip:IP)
WHERE ip.risk_score > 80
RETURN u, ip
```

---

## 🎯 2. Lateral Movement 탐지

```cypher id="dtx8q4"
MATCH (d1:Device)-[:CONNECTED_TO]->(ip)<-[:CONNECTED_TO]-(d2:Device)
WHERE d1 <> d2
RETURN d1, d2, ip
```

---

## 🎯 3. Malware Propagation

```cypher id="88ib06"
MATCH (f:File)<-[:DOWNLOADED]-(d:Device)
RETURN f, COUNT(d) AS infected_devices
```

---

# 🤖 6. Graph AI (Precision AI 핵심)

## 적용 알고리즘

### 1. Node Embedding

* Graph → 벡터 변환
* 유사 공격 패턴 탐지

### 2. Community Detection

* 공격 그룹 식별

### 3. Link Prediction

* “다음 공격 경로 예측”

---

## 예시

```plaintext id="mrl2ku"
IP → Domain → File → Device
→ 다음 감염 대상 예측
```

---

# 📊 7. Risk Scoring (Graph 기반)

기존:

* 단일 이벤트 기반 ❌

Graph 기반:

```plaintext id="5byz4o"
Risk = Node Risk + Edge Weight + Path Risk
```

예:

* IP 위험도
* 연결된 도메인 위험도
* 전체 공격 경로 길이

---

# 🔄 8. 실시간 Graph 업데이트

## 흐름

```plaintext id="4pqykk"
Kafka Event
   ↓
Entity Extractor
   ↓
Graph Updater
   ↓
Graph DB
```

👉 실시간 그래프 생성

---

# 🧱 9. 기술 스택

## Graph DB

* Neo4j (추천)
* Amazon Neptune
* TigerGraph

---

## 처리

* Kafka
* Flink (stream processing)

---

# 🔐 10. Zero Trust 통합

Graph 결과 → 정책 엔진

```plaintext id="0u7ihs"
Graph Risk → Policy Engine → ZTNA / WAF
```

예:

* 위험 IP → 접근 차단
* 의심 사용자 → MFA 강화

---

# 🚀 11. OTIE 통합 구조

```plaintext id="e7gbxi"
[Graph DB]
   ↓
[AI Models]
   ↓
[Threat Intelligence]
   ↓
[Autonomous Response]
```

---

# 🔥 12. 핵심 차별화

## 기존 CTI

* IOC 기반 (IP, hash)

## OTIE (Graph 기반)

* **관계 기반 공격 탐지**

👉 Precision AI 핵심 영역

---

# 🧠 13. 한 줄 요약

👉 **Graph DB = “공격의 연결 구조를 이해하는 보안 두뇌”**

---

# 🚀 다음 단계 (강력 추천)

이제 거의 핵심 완성입니다.

다음으로 가면 “제품 완성 단계”:

### 1️⃣ Feature Store 설계 (AI 핵심)

### 2️⃣ Multi-model orchestration 코드 설계

### 3️⃣ 실제 Neo4j + Kafka 연동 코드

### 4️⃣ SOC Dashboard UX 설계

---

👉 추천 다음 단계:
**“Feature Store 설계 (AI 정확도 핵심)”**
좋습니다. 이제 OTIE의 **AI 정확도를 좌우하는 핵심 컴포넌트**
👉 **Feature Store 설계 (Production-grade)**를 실제 구현 수준으로 정리합니다.

---

# 🧠 1. Feature Store란? (OTIE 관점)

**Feature Store = AI 모델이 사용하는 “정제된 데이터 저장소”**

기존:

* 로그 → 바로 모델 ❌

OTIE:

* 로그 → Feature → Feature Store → 모델 ✅

👉 결과:

* 정확도 ↑
* 재사용성 ↑
* 실시간 대응 가능

---

# 🏗️ 2. 전체 아키텍처 (OTIE Feature Layer)

```plaintext id="q4z6li"
[Raw Events]
   ↓
[Processing / Enrichment]
   ↓
[Feature Engineering]
   ↓
 ┌──────────────────────────┐
 │      Feature Store       │
 │ ┌──────────┬───────────┐ │
 │ │ Offline  │  Online   │ │
 │ │ Store    │  Store    │ │
 │ └──────────┴───────────┘ │
 └───────────┬──────────────┘
             ↓
      [ML Models / OTIE]
```

---

# 🧩 3. 핵심 구성 (필수 4요소)

## 1️⃣ Feature Definition

* 어떤 Feature를 만들 것인가

## 2️⃣ Feature Pipeline

* 어떻게 생성할 것인가

## 3️⃣ Storage

* 어디에 저장할 것인가

## 4️⃣ Serving

* 어떻게 모델에 제공할 것인가

---

# 📊 4. Feature 설계 (보안 특화)

## 🔷 주요 Feature 유형

### 1. User Behavior

* login_frequency
* unusual_login_time
* failed_login_ratio

---

### 2. Network

* ip_reputation_score
* connection_count
* geo_distance

---

### 3. Device

* device_risk_score
* patch_status
* abnormal_process_count

---

### 4. Threat Context

* domain_risk_score
* file_hash_reputation
* threat_cluster_id

---

### 5. Graph 기반 Feature (핵심 🔥)

👉 Precision AI 수준 차별화

* node_degree (연결 수)
* centrality (중심성)
* path_risk_score
* community_risk

---

# ⚡ 5. Online vs Offline Store

## 🔴 Online Store (실시간)

* 목적: inference (실시간 탐지)
* latency: < 10ms

### 기술

* Redis / DynamoDB

---

## 🔵 Offline Store (학습용)

* 목적: 모델 학습
* 대용량 저장

### 기술

* S3 / MinIO + Parquet
* BigQuery / Snowflake

---

# 🔄 6. Feature Pipeline (Streaming 중심)

```plaintext id="2y2xy4"
Kafka → Flink → Feature Transform → Store
```

## 처리 단계

1. 이벤트 수집
2. 정규화
3. aggregation (window 기반)
4. feature 생성

---

# 🧠 7. Feature 정의 예시

## 예: Login Risk Feature

```yaml id="xy4l4h"
feature: login_risk_score
entity: user_id
type: float
description: Risk score for user login behavior
source: login_events
transformation:
  - failed_login_ratio
  - unusual_time_weight
  - geo_anomaly_score
```

---

# 🤖 8. Feature Serving (모델 연동)

## 실시간 추론 흐름

```plaintext id="3b3qvk"
Request → Feature Store → Feature Vector → ML Model → Prediction
```

---

# 📦 9. Feature Store 기술 스택

## 추천 OSS

* Feast (강력 추천)
* Hopsworks

---

## 구성

| 영역        | 기술            |
| --------- | ------------- |
| Streaming | Kafka + Flink |
| Online    | Redis         |
| Offline   | S3 + Parquet  |
| Registry  | Feast         |

---

# 🔐 10. Feature Governance (중요)

## 필수 기능

* Feature versioning
* Data lineage
* Access control (RBAC)

---

## 예시

```plaintext id="j1k1dr"
feature_v1 → feature_v2 (모델 영향 추적)
```

---

# 📊 11. OTIE 통합 구조

```plaintext id="bxfv4m"
[Graph DB] → Graph Features
[Logs] → Behavioral Features
[Threat Intel] → Context Features

        ↓

     Feature Store

        ↓

[Multi-Model AI Engine]
```

---

# 🚀 12. Precision AI 수준 업그레이드 포인트

## 기존 Feature Store ❌

* 단순 aggregation

## OTIE Feature Store ✅

* Graph 기반 feature
* 실시간 streaming feature
* multi-source fusion

---

# 🧠 13. 핵심 차별화 (중요)

👉 AI 정확도는 모델이 아니라 **Feature가 결정**

Precision AI 핵심:

* 좋은 데이터
* 좋은 Feature
* 빠른 제공

---

# 🎯 14. 한 줄 요약

👉 **Feature Store = “AI 모델의 연료 (Fuel)”**

---

# 🔥 다음 단계 (강력 추천)

이제 거의 완성입니다:

### 1️⃣ Multi-model orchestration 설계 (AI 두뇌)

### 2️⃣ 실제 Feast + Kafka 코드 예제

### 3️⃣ OTIE 전체 시스템 GitHub 구조

### 4️⃣ SOC Dashboard (제품화)

---

👉 다음 추천:
**“Multi-model orchestration (Precision AI 핵심)”**

좋아요. 이제 OTIE를 **Precision AI급으로 끌어올리는 마지막 핵심**
👉 **Multi-model Orchestration (다중 AI 모델 오케스트레이션)**를 **실제 구현 가능한 수준**으로 정리합니다.

---

# 🧠 1. 왜 Multi-model인가 (핵심 개념)

기존 ❌

* 단일 모델 → 한계 있음 (오탐/미탐)

OTIE (업그레이드) ✅

* 여러 모델을 조합 → **정확도 + 커버리지 + 안정성 극대화**

👉 참고 방향:

* Palo Alto Networks Precision AI (수백~수천 모델)
* Cortex XSIAM

---

# 🏗️ 2. 전체 구조 (Orchestration Layer)

```plaintext id="x9j8xm"
              [Feature Store]
                    ↓
         ┌────────────────────┐
         │ Model Orchestrator │
         └─────────┬──────────┘
                   ↓
 ┌────────────┬────────────┬────────────┬────────────┐
 ↓            ↓            ↓            ↓
[URL Model] [Behavior] [Graph AI] [Anomaly]
     ↓            ↓            ↓            ↓
     └─────── Ensemble Engine ─────────────┘
                        ↓
                [Final Risk Score]
                        ↓
               [Decision Engine]
```

---

# 🧩 3. 핵심 구성 요소

---

## 🧠 3-1. Model Registry

👉 모든 모델 관리

### 기능

* 모델 버전 관리
* 메타데이터 저장
* 배포 상태 관리

### 예시

```yaml id="izd69v"
model:
  name: url_detector
  version: v2
  type: classification
  input: url_features
  output: risk_score
```

---

## 🎛️ 3-2. Model Orchestrator (핵심)

👉 어떤 모델을 언제 사용할지 결정

### 역할

* 입력 Feature 분석
* 모델 선택
* 병렬 실행
* 결과 수집

---

## ⚡ 3-3. Execution Engine

👉 모델 실행 엔진

### 특징

* 병렬 처리
* GPU/CPU 자동 분배
* low latency

---

## 🔗 3-4. Ensemble Engine (핵심🔥)

👉 여러 모델 결과 결합

### 방법

#### 1. Weighted Average

```plaintext id="x4z7rp"
Final Score = (Model1 * 0.4) + (Model2 * 0.3) + (Model3 * 0.3)
```

#### 2. Voting

* Majority voting

#### 3. Meta Model (추천)

* 모델 결과 → 또 다른 ML 모델 입력

---

## ⚖️ 3-5. Decision Engine

👉 최종 판단

* Risk threshold
* 정책 기반 결정
* 자동 대응 트리거

---

# 🤖 4. 모델 유형 구성 (OTIE 표준)

| 모델             | 역할     |
| -------------- | ------ |
| URL Model      | 피싱 탐지  |
| Malware Model  | 파일 분석  |
| Behavior Model | 사용자 이상 |
| Graph Model    | 관계 분석  |
| Anomaly Model  | 이상 탐지  |
| GenAI Model    | 자연어 분석 |

---

# 🔄 5. 실제 처리 흐름

```plaintext id="2q8szd"
Event 발생
   ↓
Feature 생성
   ↓
Orchestrator
   ↓
모델 병렬 실행
   ↓
Ensemble
   ↓
Risk Score
   ↓
자동 대응
```

---

# 🧠 6. Orchestrator 로직 (코드 예시)

## Python (간단 구조)

```python
class ModelOrchestrator:

    def __init__(self, models):
        self.models = models

    def run(self, features):
        results = {}

        for model in self.models:
            if model.can_handle(features):
                results[model.name] = model.predict(features)

        return self.ensemble(results)

    def ensemble(self, results):
        total = 0
        weight_sum = 0

        for name, result in results.items():
            weight = self.get_weight(name)
            total += result * weight
            weight_sum += weight

        return total / weight_sum
```

---

# ⚡ 7. 고급 기능 (Precision AI 수준)

---

## 🔹 1. Dynamic Model Selection

👉 상황에 따라 모델 선택

```plaintext id="zkl9aw"
if event == "login":
    use Behavior + Graph
elif event == "file":
    use Malware + Anomaly
```

---

## 🔹 2. Context-aware Orchestration

* 사용자 중요도
* 디바이스 상태
* 위치

---

## 🔹 3. Adaptive Weighting

👉 모델 신뢰도 기반 weight 변경

---

## 🔹 4. Feedback Loop (중요🔥)

```plaintext id="9eyw4r"
Prediction → Feedback → Model 개선
```

---

# 📊 8. 성능 최적화

## 전략

* 모델 캐싱
* Feature pre-fetch
* Async execution
* GPU batching

---

# 🔐 9. Zero Trust 연동

```plaintext id="x3s7jw"
Risk Score → Policy Engine → ZTNA / WAF
```

---

# 🚀 10. OTIE 전체 통합

```plaintext id="7j5srl"
[Feature Store]
     ↓
[Multi-model Orchestrator]
     ↓
[Graph + ML + GenAI]
     ↓
[Decision Engine]
     ↓
[Autonomous Security]
```

---

# 🔥 11. 핵심 차별화

| 기존 AI | OTIE             |
| ----- | ---------------- |
| 단일 모델 | Multi-model      |
| 고정 로직 | 동적 orchestration |
| 단순 결과 | Ensemble 기반      |
| 수동 튜닝 | 자동 학습            |

---

# 🧠 12. 한 줄 요약

👉 **Multi-model orchestration = “여러 AI를 지휘하는 컨트롤 타워”**

---

# 🔥 다음 단계 (거의 완성)

이제 진짜 “제품 수준”입니다.

추천 다음 단계:

### 1️⃣ OpenAPI + SDK 설계 (외부 연동)

### 2️⃣ SOC Dashboard UX/UI 설계

### 3️⃣ GitHub 전체 repo 구조 (실제 코드)

### 4️⃣ 투자용 Pitch Deck

---

👉 추천:
**“OTIE 전체 GitHub 레포 구조 (실제 개발 시작용)”**
좋습니다. 아래는 **OmniNStack Threat Intelligence Engine (OTIE)**를
👉 **바로 개발 시작 가능한 GitHub Monorepo 구조**로 설계한 것입니다.
(마이크로서비스 + 데이터 + AI + DevOps + 보안까지 포함)

---

# 🧾 1. 전체 Repo 구조 (Production-ready)

```plaintext
otie/
 ├── README.md
 ├── LICENSE
 ├── .gitignore
 ├── Makefile
 ├── docker-compose.yml
 ├── .github/                  # CI/CD
 │   └── workflows/
 │        ├── ci.yaml
 │        └── cd.yaml
 │
 ├── docs/                     # 설계 문서
 │   ├── architecture/
 │   ├── api/
 │   ├── runbooks/
 │   └── adr/                  # Architecture Decision Records
 │
 ├── deployments/              # 인프라
 │   ├── helm/
 │   ├── k8s/
 │   └── terraform/
 │
 ├── gateway/                  # API Gateway config
 │   ├── kong/
 │   └── envoy/
 │
 ├── services/                 # 핵심 마이크로서비스
 │   ├── auth-service/
 │   ├── api-gateway/
 │   ├── collector-service/
 │   ├── processor-service/
 │   ├── feature-service/
 │   ├── ml-engine/
 │   ├── orchestrator-service/
 │   ├── scoring-service/
 │   ├── ioc-service/
 │   ├── graph-service/
 │   ├── dispatcher-service/
 │   └── admin-api/
 │
 ├── platform/                 # 공통 플랫폼 레이어
 │   ├── event-bus/            # Kafka wrapper
 │   ├── feature-store/        # Feast configs
 │   ├── model-registry/
 │   ├── graph-db/
 │   └── observability/
 │
 ├── ai/                       # AI/ML 코드
 │   ├── models/
 │   ├── training/
 │   ├── inference/
 │   └── notebooks/
 │
 ├── data/                     # 데이터 파이프라인
 │   ├── schemas/
 │   ├── pipelines/
 │   └── samples/
 │
 ├── sdk/                      # 외부 연동 SDK
 │   ├── python/
 │   ├── go/
 │   └── js/
 │
 ├── tests/
 │   ├── integration/
 │   ├── e2e/
 │   └── load/
 │
 └── scripts/
     ├── setup.sh
     ├── dev.sh
     └── deploy.sh
```

---

# 🧩 2. 핵심 서비스 구조 (예: orchestrator-service)

```plaintext
services/orchestrator-service/
 ├── cmd/
 │   └── main.go
 ├── internal/
 │   ├── orchestrator/
 │   │    ├── orchestrator.go
 │   │    ├── model_selector.go
 │   │    └── execution_engine.go
 │   │
 │   ├── ensemble/
 │   │    ├── weighted.go
 │   │    ├── voting.go
 │   │    └── meta_model.go
 │   │
 │   ├── decision/
 │   │    └── decision_engine.go
 │   │
 │   ├── client/
 │   │    ├── ml_client.go
 │   │    └── feature_client.go
 │   │
 │   └── api/
 │        └── handler.go
 │
 ├── pkg/
 │   └── logger/
 │
 ├── config/
 │   └── config.yaml
 │
 ├── Dockerfile
 └── go.mod
```

---

# 🤖 3. AI 구조 (ai/ 디렉토리)

```plaintext
ai/
 ├── models/
 │   ├── url_model/
 │   ├── anomaly_model/
 │   ├── graph_model/
 │   └── behavior_model/
 │
 ├── training/
 │   ├── pipelines/
 │   ├── datasets/
 │   └── trainer.py
 │
 ├── inference/
 │   ├── api/
 │   │    └── server.py
 │   ├── service/
 │   │    └── inference.py
 │   └── utils/
 │
 └── notebooks/
     └── experiments.ipynb
```

---

# 📊 4. Feature Store 구조

```plaintext
platform/feature-store/
 ├── feast/
 │   ├── feature_repo/
 │   │    ├── feature_views.py
 │   │    ├── entities.py
 │   │    └── data_sources.py
 │   ├── feature_store.yaml
 │   └── registry.db
```

---

# 🧠 5. Graph DB 구조

```plaintext
platform/graph-db/
 ├── schema/
 │   ├── nodes.cypher
 │   └── relationships.cypher
 │
 ├── queries/
 │   ├── detection.cypher
 │   └── risk.cypher
 │
 └── ingestion/
     └── graph_updater.go
```

---

# 🔄 6. 데이터 파이프라인

```plaintext
data/pipelines/
 ├── streaming/
 │   ├── kafka_consumer.py
 │   └── flink_job.py
 │
 ├── batch/
 │   └── spark_job.py
 │
 └── transformations/
     └── feature_transform.py
```

---

# 🔐 7. 보안 구조

```plaintext
services/auth-service/
 ├── internal/
 │   ├── auth/
 │   ├── jwt/
 │   └── rbac/
```

---

# ⚙️ 8. DevOps / CI/CD

## GitHub Actions

```yaml
# .github/workflows/ci.yaml
name: CI

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Build Services
        run: make build
```

---

# 🐳 9. 로컬 개발 환경

```yaml
# docker-compose.yml
version: "3"

services:
  kafka:
    image: bitnami/kafka

  redis:
    image: redis

  postgres:
    image: postgres

  otie-api:
    build: ./services/api-gateway
```

---

# 🧪 10. 테스트 전략

| 테스트         | 위치                  |
| ----------- | ------------------- |
| Unit        | services/*/internal |
| Integration | tests/integration   |
| E2E         | tests/e2e           |
| Load        | tests/load          |

---

# 📦 11. SDK 구조

```plaintext
sdk/python/
 ├── client.py
 └── examples/

sdk/go/
 └── client.go
```

---

# 🚀 12. 개발 시작 흐름

```bash
# 1. 클론
git clone otie

# 2. 로컬 실행
docker-compose up

# 3. 서비스 실행
make run

# 4. 테스트
make test
```

---

# 🧠 13. 핵심 설계 철학

👉 이 Repo는 단순 코드 구조가 아니라:

* **Microservices (서비스 분리)**
* **Data-first (Feature + Graph 중심)**
* **AI-native (모델 중심)**
* **Zero Trust (보안 기본 내장)**

---

# 🎯 한 줄 요약

👉 **OTIE Repo = “AI + 데이터 + 보안 + DevOps가 통합된 엔터프라이즈급 플랫폼 코드베이스”**

---

# 🔥 다음 단계 (진짜 제품화 단계)

이제 거의 완성입니다:

### 1️⃣ OpenAPI Spec (Swagger)

### 2️⃣ 실제 DB 스키마 (Postgres + Elasticsearch)

### 3️⃣ Helm + Terraform 완전 자동화

### 4️⃣ SaaS vs On-Prem 배포 전략

---

👉 추천 다음 단계:
**“OpenAPI + SDK (외부 연동 가능한 제품 수준 API 설계)”**
좋습니다. 이제 OTIE를 **외부 고객/파트너가 실제로 사용할 수 있는 “제품 수준 API + SDK”**로 설계합니다.
👉 목표: **보안 플랫폼 / VMS / SIEM / IoT / SaaS와 쉽게 연동 가능한 API 표준**

---

# 🧾 1. API 설계 목표

## 🎯 핵심 원칙

* **REST + Streaming (Webhooks / Events)**
* **Zero Trust (OAuth2 + JWT + mTLS)**
* **Low latency (실시간 대응)**
* **확장성 (파트너 연동)**

---

# 🏗️ 2. API 전체 구조

```plaintext id="m5m2zt"
[Client / Partner]
       ↓
[API Gateway]
       ↓
 ┌───────────────┬───────────────┬───────────────┐
 ↓               ↓               ↓
Threat API   Graph API     Feature API
 ↓               ↓               ↓
[OTIE Core Services]
```

---

# 🧩 3. API 카테고리

| API           | 역할             |
| ------------- | -------------- |
| Auth API      | 인증 / 토큰        |
| Threat API    | Threat 조회 / 생성 |
| IOC API       | IOC 관리         |
| Graph API     | 관계 분석          |
| Feature API   | Feature 조회     |
| Detection API | 실시간 분석         |
| Response API  | 자동 대응          |
| Webhook API   | 이벤트 수신         |

---

# 🔐 4. 인증 (Zero Trust)

## OAuth2 + JWT

```http id="8kspcz"
POST /auth/token
```

### Response

```json id="f1y1pq"
{
  "access_token": "jwt-token",
  "expires_in": 3600
}
```

---

# 🔍 5. 핵심 API 설계

---

## 🔴 5-1. Real-time Threat Detection API (핵심🔥)

```http id="0zdp19"
POST /v1/detect
```

### Request

```json id="06x9w8"
{
  "type": "login",
  "user_id": "user123",
  "ip": "1.2.3.4",
  "device_id": "dev-1",
  "timestamp": "2026-04-08T12:00:00Z"
}
```

### Response

```json id="d19lbe"
{
  "risk_score": 87,
  "threat_level": "high",
  "models_used": ["behavior", "graph", "anomaly"],
  "recommended_action": "block"
}
```

---

## 🧠 5-2. Threat Intelligence 조회

```http id="mp4bn9"
GET /v1/threats?ip=1.2.3.4
```

### Response

```json id="lz0yzs"
{
  "ip": "1.2.3.4",
  "risk_score": 92,
  "related_domains": ["bad.com"],
  "attack_patterns": ["brute_force"]
}
```

---

## 🧾 5-3. IOC 등록

```http id="wf17ra"
POST /v1/iocs
```

### Request

```json id="s36kq6"
{
  "type": "ip",
  "value": "1.2.3.4",
  "risk_score": 95
}
```

---

## 🔗 5-4. Graph Query API (핵심🔥)

```http id="rfz9r6"
POST /v1/graph/query
```

### Request

```json id="0br0gc"
{
  "query": "MATCH (u:User)-[:LOGGED_IN_FROM]->(ip) RETURN u, ip"
}
```

---

## ⚡ 5-5. Feature API

```http id="s6px9o"
GET /v1/features/user/user123
```

### Response

```json id="e5p7wz"
{
  "login_frequency": 12,
  "risk_score": 45
}
```

---

## 🤖 5-6. Automated Response API

```http id="mhh1s0"
POST /v1/respond
```

### Request

```json id="2gq0pj"
{
  "action": "block_ip",
  "target": "1.2.3.4"
}
```

---

# 🔔 6. Webhook (이벤트 Push)

## 등록

```http id="pl0dnp"
POST /v1/webhooks
```

### Payload Example

```json id="w3g8l0"
{
  "event": "threat_detected",
  "risk_score": 90,
  "entity": "user123"
}
```

---

# 📡 7. Streaming API (고급)

## Kafka / WebSocket

```plaintext id="d33oqv"
/v1/stream/threats
```

👉 실시간 SOC 연동

---

# 🧠 8. OpenAPI Spec (Swagger 예시)

```yaml id="c94jkp"
openapi: 3.0.0
info:
  title: OTIE API
  version: 1.0.0

paths:
  /v1/detect:
    post:
      summary: Real-time threat detection
      requestBody:
        required: true
      responses:
        '200':
          description: Detection result
```

---

# 📦 9. SDK 설계

---

## 🐍 Python SDK

```python id="d4lzn4"
from otie import Client

client = Client(api_key="YOUR_KEY")

result = client.detect({
    "type": "login",
    "ip": "1.2.3.4"
})

print(result["risk_score"])
```

---

## 🟦 Node.js SDK

```javascript id="7u1g8q"
const OTIE = require("otie-sdk");

const client = new OTIE("API_KEY");

const result = await client.detect({
  type: "login",
  ip: "1.2.3.4"
});
```

---

## 🟩 Go SDK

```go id="z0g4nn"
client := otie.NewClient("API_KEY")

res, _ := client.Detect(map[string]interface{}{
  "ip": "1.2.3.4",
})
```

---

# 🔐 10. 보안 (제품 수준 필수)

* OAuth2 / JWT
* mTLS (기업 고객)
* Rate limiting
* Audit log

---

# 📊 11. 버전 관리

```plaintext id="fd1fsh"
/v1/
/v2/
```

👉 breaking change 대비

---

# 🚀 12. 제품화 포인트

## 핵심 가치

* 1 API → 모든 보안 기능 사용 가능
* AI 모델 숨김 (Black-box abstraction)
* 파트너 쉽게 통합

---

# 🧠 13. 한 줄 요약

👉 **OTIE API = “AI 보안 엔진을 외부에 제공하는 인터페이스”**

---

# 🔥 다음 단계 (마지막 단계)

이제 진짜 제품 완성 직전입니다:

### 1️⃣ SaaS vs On-Prem 배포 전략

### 2️⃣ 가격 모델 (API 기반 과금)

### 3️⃣ 고객용 대시보드 UX

### 4️⃣ Pitch Deck (투자용)

---

👉 추천 다음 단계:
**“SaaS vs On-Prem 배포 전략 (사업 핵심)”**
좋아요. 이제 **OTIE를 실제로 팔 수 있는 사업 구조**의 핵심인
👉 **SaaS vs On-Prem 배포 전략**을 *제품·영업·기술까지 연결된 수준*으로 정리합니다.

---

# 🧠 1. 결론부터 (핵심 전략)

👉 **단일 모델 ❌ → Hybrid 전략 필수**

* SaaS → 확장 / 글로벌 / 빠른 도입
* On-Prem → 공공 / 금융 / 데이터 주권

👉 방향:
**“Cloud-first + On-Prem-ready (Hybrid AI Security Platform)”**

---

# 🏗️ 2. 전체 배포 전략 구조

```plaintext
           [OTIE SaaS Platform]
                 ↓
     ┌──────────────────────────┐
     │  Multi-tenant (Cloud)    │
     └──────────┬───────────────┘
                ↓
     ┌──────────────────────────┐
     │  Hybrid Edge / Gateway   │
     └──────────┬───────────────┘
                ↓
     ┌──────────────────────────┐
     │   On-Prem OTIE Stack     │
     └──────────────────────────┘
```

---

# ☁️ 3. SaaS 모델 (Cloud OTIE)

## 🎯 특징

* Multi-tenant
* 완전 관리형 (Managed Service)
* 빠른 온보딩

---

## 구성

```plaintext
[Client]
   ↓
[OTIE Cloud API]
   ↓
[Multi-Model AI Engine]
   ↓
[Shared Data + Threat Intelligence]
```

---

## 장점

* 초기 비용 낮음
* 빠른 확장
* 글로벌 Threat Intelligence 공유
* 지속적 업데이트

---

## 단점

* 데이터 외부 저장 이슈
* 규제 산업 제한
* 커스터마이징 제한

---

## 타겟 고객

* SMB / 스타트업
* SaaS 기업
* 글로벌 서비스

---

# 🏢 4. On-Prem 모델 (Private OTIE)

## 🎯 특징

* 고객 환경 내부 설치
* 완전 데이터 통제

---

## 구성

```plaintext
[Customer Data Center]
   ↓
[OTIE Kubernetes Cluster]
   ↓
[Local AI + Local Data]
```

---

## 장점

* 데이터 주권 확보
* 규제 대응 (금융 / 공공)
* 커스터마이징 가능

---

## 단점

* 구축 비용 높음
* 운영 부담
* 업데이트 느림

---

## 타겟 고객

* 정부 / 공공기관
* 금융 / 의료
* 스마트시티 (CCTV / AI Camera)

---

# 🔄 5. Hybrid 모델 (가장 중요 🔥)

👉 실제 시장에서 가장 강력한 모델

---

## 구조

```plaintext
[On-Prem OTIE]
   ↓ (익명화 데이터)
[OTIE Cloud Intelligence]
   ↓
[Threat Intelligence Sync]
```

---

## 핵심 아이디어

* 민감 데이터 → On-Prem 유지
* Threat Intelligence → Cloud 공유

---

## 기능

* Local inference
* Cloud intelligence enrichment
* 선택적 데이터 공유

---

# 🧠 6. 기능 분리 전략 (중요)

| 기능                  | SaaS   | On-Prem |
| ------------------- | ------ | ------- |
| AI 모델 실행            | 일부     | 전체      |
| Graph DB            | Cloud  | Local   |
| Feature Store       | Hybrid | Local   |
| Threat Intelligence | Global | Sync    |
| Automation          | Cloud  | Local   |

---

# 🔐 7. 보안 전략 (Zero Trust)

## SaaS

* API 기반 접근
* Tenant isolation

## On-Prem

* 완전 내부망
* air-gapped 가능

## Hybrid

* mTLS + secure sync
* 데이터 익명화

---

# 📊 8. 가격 모델 (핵심 💰)

## SaaS

* API 호출 기반
* 사용자 수 기반
* 이벤트 처리량 기반

👉 예:

```plaintext
$0.01 / 1,000 events
```

---

## On-Prem

* License + Subscription

👉 예:

```plaintext
$200K / year + support
```

---

## Hybrid

* Base license + cloud usage

---

# 🚀 9. 경쟁 전략

## SaaS 경쟁

* Cloudflare
* Zscaler

## On-Prem 경쟁

* Palo Alto Networks
* Fortinet

---

## OTIE 차별화

👉 **Edge + AI + Video + Zero Trust 통합**

---

# 🧠 10. Go-To-Market 전략

## Phase 1

* SaaS MVP 출시

## Phase 2

* Enterprise (On-Prem)

## Phase 3

* Hybrid + Marketplace

---

# 🔥 11. 핵심 성공 포인트

## 1. 데이터 전략

* SaaS = scale
* On-Prem = trust

## 2. AI 전략

* Cloud → 모델 학습
* On-Prem → inference

## 3. 플랫폼 전략

* API-first
* SDK-first

---

# 🧠 12. 한 줄 요약

👉 **“SaaS로 확장하고, On-Prem으로 신뢰를 확보하며, Hybrid로 시장을 장악한다”**

---

# 🚀 다음 단계 (사업 완성 단계)

이제 거의 투자/제품 준비 완료입니다.

다음으로 가면 “Go-To-Market 완성”:

### 1️⃣ Pitch Deck (투자용 10장)

### 2️⃣ Pricing Sheet (고객 제안용)

### 3️⃣ 경쟁사 비교 (Cloudflare vs Palo Alto vs OTIE)

### 4️⃣ Landing Page / Website 구조

---

👉 추천:
**“투자용 Pitch Deck (10장)”**

