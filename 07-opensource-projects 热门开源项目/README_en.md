**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 07-opensource-projects: Popular Open Source Projects

## Module Overview

This chapter is a curated resource list instead of a demo collection. It highlights Go open source projects that are worth following over the long term.

As of `2026-03-30`, this list prioritizes projects that are:

- consistently influential in the Go ecosystem
- valuable in real engineering work
- strong in architecture, source layout, or documentation quality
- useful for learners moving from syntax practice to reading production code

This is not a strict star ranking. It is a learning-oriented recommendation list.

## Suggested Reading Order

- Beginner-friendly: `Gin`, `Chi`, `Cobra`, `Hugo`
- Engineering growth: `Caddy`, `Traefik`, `Prometheus`, `MinIO`
- Architecture growth: `Kubernetes`, `etcd`, `containerd`, `NATS`
- Frontier topics: `Milvus`, `Loki`, `Ollama`, `go-ethereum`

## Project List

| Category | Project | URL | Short Comment | Why It Is Recommended |
| ---- | ---- | ---- | ---- | ---- |
| Web Framework | Gin | [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) | One of the most widely used Go web frameworks. | Great for learning routing, middleware, request binding, and JSON API structure in a practical style. |
| Web Router | Chi | [github.com/go-chi/chi](https://github.com/go-chi/chi) | A lightweight router with strong standard-library flavor. | Helps you understand idiomatic `net/http` composition instead of relying only on full-stack frameworks. |
| CLI Tooling | Cobra | [github.com/spf13/cobra](https://github.com/spf13/cobra) | A de facto standard for Go CLI applications. | Excellent for learning subcommand structure, flags, help output, and large CLI organization. |
| Static Site Generator | Hugo | [github.com/gohugoio/hugo](https://github.com/gohugoio/hugo) | A highly mature static site generator. | A strong case study for large CLI apps, filesystem-heavy workflows, templating, and performance tuning. |
| Web Server | Caddy | [github.com/caddyserver/caddy](https://github.com/caddyserver/caddy) | A modern web server known for automatic HTTPS. | Great for learning modular design, config-driven architecture, plugins, and production-oriented server engineering. |
| Cloud Native | Kubernetes | [github.com/kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) | One of the most influential Go projects in cloud native systems. | A classic study target for large repository organization, controller patterns, and declarative system design. |
| Distributed Storage | etcd | [github.com/etcd-io/etcd](https://github.com/etcd-io/etcd) | A compact but hard-core distributed key-value store. | Ideal for learning Raft, consistency, leases, watch mechanics, and distributed systems in Go. |
| Observability | Prometheus | [github.com/prometheus/prometheus](https://github.com/prometheus/prometheus) | A core project in the monitoring world. | Useful for learning metrics collection, service discovery, rule evaluation, and time-series engineering. |
| Gateway / Proxy | Traefik | [github.com/traefik/traefik](https://github.com/traefik/traefik) | A common cloud-native reverse proxy and gateway. | Worth reading for dynamic config handling, provider abstraction, middleware pipelines, and traffic entry design. |
| Container Runtime | containerd | [github.com/containerd/containerd](https://github.com/containerd/containerd) | A core component in the container stack. | Helps you understand runtime internals, daemon design, task lifecycle management, and cloud-native foundations. |
| Messaging | NATS Server | [github.com/nats-io/nats-server](https://github.com/nats-io/nats-server) | A fast and elegant messaging server. | Good for learning protocol implementation, concurrent networking, pub/sub patterns, and high-performance service code. |
| Object Storage | MinIO | [github.com/minio/minio](https://github.com/minio/minio) | A star project in S3-compatible object storage. | Strong for studying API compatibility, storage architecture, and real production engineering patterns. |
| Vector Database | Milvus | [github.com/milvus-io/milvus](https://github.com/milvus-io/milvus) | A representative vector database in the AI era. | Useful if you want to connect Go learning with AI infrastructure and large-scale retrieval systems. |
| Logging System | Loki | [github.com/grafana/loki](https://github.com/grafana/loki) | A central logging project in the Grafana ecosystem. | Good for studying log indexing tradeoffs, distributed components, and observability system design. |
| LLM Tooling | Ollama | [github.com/ollama/ollama](https://github.com/ollama/ollama) | A popular local LLM runtime tool. | Shows how Go is used in AI tooling for model management, local serving, APIs, and cross-platform packaging. |
| Blockchain | go-ethereum | [github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) | One of the best-known Go blockchain projects. | A challenging but rewarding codebase for studying P2P networking, state sync, transaction pools, and VM-related engineering. |

## How To Read These Repositories

- Do not start by reading every file in order. Start with the README, directory layout, and entry points.
- Read `main`, config loading, route registration, and module initialization before diving into deep implementation details.
- Pick repositories that match your current goal: Web learners should start with Gin or Chi; distributed systems learners should start with etcd or NATS.
- Keep notes while reading, especially about layering, shared abstractions, configuration flow, and runtime state flow.

## Focus Areas

### 1. Learning Go Web Engineering

- Gin
- Chi
- Caddy
- Traefik

### 2. Learning CLI and Tool Development

- Cobra
- Hugo
- Ollama

### 3. Learning Cloud Native and Distributed Systems

- Kubernetes
- etcd
- containerd
- NATS
- Prometheus

### 4. Learning Storage, Search, and Data Infrastructure

- MinIO
- Milvus
- Loki

## Extra Advice

- If this is your first time reading large Go repositories, start with `Cobra`, `Gin`, `Chi`, or `Hugo`.
- If you have finished the basics, core features, standard library, and framework chapters in this repository, you will get much more value from `Kubernetes`, `etcd`, and `go-ethereum`.
- Growth does not come from “having seen many repositories”. It comes from recognizing the repeated patterns behind structure, abstraction, concurrency, and error handling.
