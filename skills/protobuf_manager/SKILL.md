# Skill: OTIE Protobuf Manager

## Overview

This skill provides instructions and tools for managing gRPC/Protobuf definitions across the OmniNStack-OTIE platform.

## Instructions

1.  Store all `.proto` files in the source of truth directory: `./api/v1/`.
2.  Ensure every service has a unique gRPC package name following the pattern `otie.{service_name}.v1`.
3.  Use `protoc` or `buf` for compilation into both Go (server-side) and Python (client-side for ML-Engine).
4.  Always generate both a server stub and a client library for each service definition.

## Recommended Tools

-   `buf` CLI for linting and breaking change detection.
-   `protoc-gen-go-grpc` for Go server implementation.
-   `protoc-gen-python` for Python client implementation.

## Example Flow

```bash
# Linting
buf lint

# Generation
buf generate
```
