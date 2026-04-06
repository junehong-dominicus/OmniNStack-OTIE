# OTIE Workflow Review Guide

This document provides a high-level summary of the development and review workflows for the **OmniNStack Threat Intelligence Engine (OTIE)**. It is designed to help architects, leads, and developers understand the project lifecycle and the critical points where human feedback is required.

## 1. Workflow Map

The OTIE development process follows a four-phase lifecycle:

| Phase | Workflow | Key Outcome |
| :--- | :--- | :--- |
| **Phase 1: Setup** | `init_project_structure` | Baseline project structure and shared resources. |
| **Phase 2: Dev** | `add_new_microservice` | New functional service with verified contracts. |
| **Phase 3: Test** | `setup_local_dev_env` | Stable local execution environment. |
| **Phase 4: Verify** | `pipeline_e2e_verification` | Verified end-to-end data pipeline integrity. |
| **Phase 5: Deploy** | `ci_cd_github_actions` | Automated and secure production deployment. |

---

## 2. Critical Review & Feedback Points

To maintain the project's **Zero Trust** and **High Performance** standards, "Review & Feedback" steps are integrated into every workflow:

### A. Project Infrastructure (`init_project_structure`)

- **Review Point**: After directory creation and Go module initialization.
- **Reviewer**: Project Lead / Architect.
- **Focus**: Logical separation of `cmd`, `internal`, and `pkg` folders.

### B. Service Scaffolding (`add_new_microservice`)

- **Review Point**: After gRPC definitions and Go logic bootstrapping.
- **Reviewer**: Project Architect / Senior Developer.
- **Focus**: Adherence to the interface contract and Kafka producer/consumer patterns.

### C. Environment Stability (`setup_local_dev_env`)

- **Review Point**: After launching the Docker Compose stack.
- **Reviewer**: DevOps Engineer / Developer.
- **Focus**: Verification of healthy heartbeat for Kafka, PostgreSQL, and Elasticsearch.

### D. Pipeline Integrity (`pipeline_e2e_verification`)

- **Review Point**: After end-to-end data injection.
- **Reviewer**: SecOps / ML Engineer.
- **Focus**: Accuracy of the risk score evaluation and correct IOC storage in the DB.

### E. Pipeline Security (`ci_cd_github_actions`)

- **Review Point**: After YAML configuration of GitHub Actions.
- **Reviewer**: Security Auditor / DevOps.
- **Focus**: Ensuring no secret leakage and adherence to build security standards.

---

## 3. How to Use Workflows

Each workflow is stored as a Markdown file in the `.agent/workflows/` directory. They are designed to be followed sequentially by agentic assistants or human developers.

> [!TIP]
> Use the **`architect.agent`** for strategic design reviews and the **`devops_engineer.agent`** for validating CI/CD and Helm configurations.
