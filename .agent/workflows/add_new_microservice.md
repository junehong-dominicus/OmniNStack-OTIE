# Workflow: Add New Microservice

---
description: How to add a new Go-based microservice to the OTIE project.
---

1. Define the service interface in `api/v1/{service}.proto`.

2. Compile the Protobuf:

   ```bash
   # Using 'buf'
   buf generate
   ```

3. Create the directory structure:

   ```bash
   mkdir cmd/{service}
   mkdir internal/{service}
   mkdir internal/{service}/handler
   mkdir internal/{service}/service
   mkdir internal/{service}/repository
   ```

4. Create the `main.go` entry point in `cmd/{service}/` and use the `service_scaffolder` skill to bootstrap common logic (e.g., gRPC server, Kafka producer/consumer).

5. **Review & Feedback**: Present the scaffolded service structure and `main.go` to the project architect for logic and interface verification.

6. Implement context-aware business logic in `internal/{service}/service/`.

7. Add the service to the `docker-compose.yaml` file for local development.

8. Create the Helm chart in `deployments/helm/{service}/` using the existing templates as a base.

9. Update the `.github/workflows/ci.yaml` to include the building and testing of the new service.
