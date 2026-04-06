# Workflow: CI/CD with GitHub Actions

---
description: Standards for GitHub Actions workflow implementation.
---

1. Create a `.github/workflows/` directory in the repository.

2. Create a `ci.yaml` file to automate testing for Go and Python microservices.

   ```yaml
   name: CI for OTIE Platform
   on: [push, pull_request]

   jobs:
     test-go:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v4
         - name: Set up Go
           uses: actions/setup-go@v5
           with:
             go-version: '1.22'
         - run: go test ./...

     test-python:
       runs-on: ubuntu-latest
       steps:
         - uses: actions/checkout@v4
         - name: Set up Python
           uses: actions/setup-python@v5
           with:
             python-version: '3.11'
         - run: pip install -r ml_engine/requirements.txt && pytest ml_engine/
   ```

3. **Review & Feedback**: Perform a security audit of the GitHub Actions YAML configuration. Ensure that no sensitive information is exposed and that the build process follows project standards.

4. Configure secret management in the GitHub repository (e.g., Docker registry credentials, K8s secrets).

5. Add a `deploy.yaml` for automated deployment to the OTIE staging environment upon merge to the `main` branch.

6. Use the `github_actions_v1` agent to monitor and debug workflow failures.
