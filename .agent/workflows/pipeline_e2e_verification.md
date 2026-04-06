# Workflow: Pipeline E2E Verification

---
description: Workflow for testing the full "Collector -> Kafka -> ML -> IOC" flow.
---

1. Start all core services in the local Docker environment (via `setup_local_dev_env`).

2. Use the `test_event_generator` tool to inject a sample malicious log into the `collector-service` endpoint.

3. Verify that the event is published to the Kafka topic.

4. Verify that the `processor-service` consumes and enriches the event.

5. Monitor the `ml_engine-service` logs to ensure inference is performed on the event features.

6. Confirm that the `scoring-service` assigns a high risk score to the event.

7. **Review & Feedback**: Check the aggregate results of the pipeline and confirm that the risk evaluation is correct based on the input scenario.

8. Check the `ioc_service-service` PostgreSQL database to ensure a new threat entry is recorded.

9. Verify that the `dispatcher-service` sends a webhook alert (Mock Dispatcher endpoint).

10. Use the `otie_verify_test_logs` script to confirm completion of the end-to-end pipeline.
