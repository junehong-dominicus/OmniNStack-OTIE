# Skill: OTIE Security Audit

## Overview

This skill provides instructions for auditing the OTIE platform according to the Zero Trust principles.

## Audit Checklist

- **mTLS**: Verify that `istio-proxy` or equivalent mTLS layer is active between all project namespaces.
- **Service Identity**: Check SPIFFE SVID for pod-to-pod identity propagation.
- **JWT**: Verify short-lived tokens are required for all non-gRPC internal communication (e.g., via Admin UI API).
- **Network Policy**: Validate that `DefaultDenyAll` is applied and specific egress/ingress rules exist to allow only Kafka/DB traffic.
- **Configuration Security**: Verify no secrets are stored globally in `values.yaml` and all are pulled from HashiCorp Vault.

## Tools

- `istioctl proxy-config` for mTLS check.
- `kube-bench` for Kubernetes security policy auditing.
- `snyk` or `trivy` for container image vulnerabilities.
