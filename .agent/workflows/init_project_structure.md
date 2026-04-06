# Workflow: Initial Project Structure Setup

---
description: How to initialize the OTIE project structure from scratch.
---

1. Create the root directory structure for the OTIE platform.

   ```bash
   mkdir cmd internal pkg deployments scripts api .github .github/workflows .agent .agent/workflows skills
   ```

2. Initialize the Go module and shared utility package.

   ```bash
   go mod init github.com/OmniNStack/OmniNStack-OTIE
   ```

3. **Review & Feedback**: Present the initialized folder structure to the project lead for approval before proceeding with core infrastructure configuration.

4. Establish the API definition folder `api/v1` for Protobuf files.

5. Create the initial `docker-compose.yaml` in the root for Kafka and PostgreSQL setup.
