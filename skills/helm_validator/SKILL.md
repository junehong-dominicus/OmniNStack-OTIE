# Skill: OTIE Helm Validator

## Overview

This skill provides instructions for linting and validating OTIE-specific Helm charts for the platform.

## Validation Steps

1.  Run `helm lint deployments/helm/{service_name}` for basic syntax checking.
2.  Ensure `values.yaml` defines `images.imagePullPolicy: IfNotPresent` for dev/staging.
3.  Validate service account and network policy definitions:
    -   Ensure a service account exists for each deployment.
    -   Restrict Ingress/Egress traffic between services to only necessary pods.
4.  Run `kubescore` for security best practices.
    -   Resource limits and requests must be defined for each container.
    -   Probes (Liveness/Readiness) must be implemented for all HTTP/gRPC services.

## Automated Checks

```bash
# Example helm template for dry-run
helm template otie deployments/helm/{service_name} --values values.yaml
```
