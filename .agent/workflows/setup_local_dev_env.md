# Workflow: Setup Local Dev Environment

---
description: Setup local Kafka and PostgreSQL for OTIE development and testing.
---

1. Ensure Docker and Docker Compose are installed and running.

2. Create the `docker-compose.yaml` in the root:

   ```yaml
   version: '3.8'
   services:
     kafka:
       image: bitnami/kafka:latest
       ports:
         - "9092:9092"
       environment:
         - KAFKA_CFG_NODE_ID=1
         - KAFKA_CFG_PROCESS_ROLES=controller,broker
         - KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=1@kafka:9093
         - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093
         - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
         - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT
         - KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
         - KAFKA_CFG_INTER_BROKER_LISTENER_NAME=PLAINTEXT

     postgres:
       image: postgres:latest
       ports:
         - "5432:5432"
       environment:
         - POSTGRES_USER=otie_user
         - POSTGRES_PASSWORD=otie_password
         - POSTGRES_DB=otie_db

     elasticsearch:
       image: elasticsearch:8.12.0
       ports:
         - "9200:9200"
       environment:
         - discovery.type=single-node
         - xpack.security.enabled=false
   ```

3. Start the services:

   ```bash
   docker-compose up -d
   ```

4. **Review & Feedback**: Verify that all containers are running and healthy. Ensure that connectivity to Kafka and PostgreSQL is successfully established before proceeding with local development.

5. Verify service status and connectivity using the `otie_test_conn` tool.
