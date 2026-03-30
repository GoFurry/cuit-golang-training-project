**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 04-go-frameworks: Mainstream Go Frameworks

## Module Overview

This chapter focuses on the Go frameworks and toolchains that are most useful for day-one learning. Instead of building oversized production scaffolds, it provides a set of runnable, comparable, teaching-friendly demos.

Every demo that depends on third-party packages is organized as an independent Go module and locked with:

```go
go 1.26
toolchain go1.26.0
```

This keeps dependencies isolated and makes each example easy to run on its own.

## Directory Layout

```text
04-go-frameworks Mainstream Frameworks/
├── 01-web-frameworks Web Frameworks/
│   ├── 001-demo-gin
│   ├── 002-demo-echo
│   ├── 003-demo-fiber
│   └── 004-demo-chi
├── 02-microservices Microservice Frameworks/
│   └── 001-demo-grpc
├── 03-high-performance-http High-Performance HTTP/
│   └── 001-demo-fasthttp
├── 04-cli-frameworks CLI Frameworks/
│   └── 001-demo-cobra
├── 05-gui-frameworks GUI Frameworks/
│   └── 001-demo-wails
├── 06-crawler-frameworks Crawler Frameworks/
│   └── 001-demo-colly
├── 07-orm-frameworks ORM Frameworks/
│   ├── 001-demo-gorm
│   └── 002-demo-xorm
└── 08-testing-frameworks Testing Frameworks/
    ├── 001-demo-testify
    └── 002-demo-goconvey
```

## Demo List

### 01-web-frameworks Web Frameworks

- `001-demo-gin`: routing, middleware, JSON binding, API responses
- `002-demo-echo`: context-based handlers, parameter parsing, error handling
- `003-demo-fiber`: Express-style API and fast HTTP handling
- `004-demo-chi`: lightweight router with standard-library-friendly middleware

### 02-microservices Microservice Frameworks

- `001-demo-grpc`: define services with `proto3` and run client/server code generated from protobuf

### 03-high-performance-http High-Performance HTTP

- `001-demo-fasthttp`: learn a lower-level high-performance HTTP handler style

### 04-cli-frameworks CLI Frameworks

- `001-demo-cobra`: multi-command CLI app with flags and help output

### 05-gui-frameworks GUI Frameworks

- `001-demo-wails`: desktop app development using Go plus a frontend stack

### 06-crawler-frameworks Crawler Frameworks

- `001-demo-colly`: scrape a local test website to learn selectors, links, and extraction

### 07-orm-frameworks ORM Frameworks

- `001-demo-gorm`: model definition, migration, insert, query, update
- `002-demo-xorm`: compare another ORM style built around engine/session usage

### 08-testing-frameworks Testing Frameworks

- `001-demo-testify`: assertions, error checks, and test structure
- `002-demo-goconvey`: BDD-style testing syntax

## How to Run

### 1. Run Web / HTTP / CLI / ORM / Crawler demos

```bash
cd "01-web-frameworks Web框架/001-demo-gin"
go mod tidy
go run .
```

### 2. Run the gRPC demo

Generate protobuf code first, then start the server:

```bash
cd "02-microservices 微服务框架/001-demo-grpc"
go generate ./...
go mod tidy
go run ./server
```

In another terminal:

```bash
cd "02-microservices 微服务框架/001-demo-grpc"
go run ./client
```

### 3. Run testing demos

```bash
cd "08-testing-frameworks 测试框架/001-demo-testify"
go mod tidy
go test ./...
```

### 4. Run the Wails demo

```bash
cd "05-gui-frameworks GUI框架/001-demo-wails"
go mod tidy
wails dev
```

## Learning Path

- Start with `Gin`, then compare it with `Echo`, `Fiber`, and `Chi`.
- Learn `gRPC` before moving to engineering-heavy frameworks like Kratos, Kitex, or go-zero.
- Practice ORM examples together with SQLite or MySQL so the abstraction makes sense.
- Use testing frameworks together with the standard `testing` package instead of treating them as magic.
- Treat Wails as a practical first step into GUI development from the Go ecosystem.

## Suggested Next Steps

This chapter intentionally prioritizes lightweight and teachable demos. After finishing it, you can extend the repository with:

- Web: `Beego`, `Hertz`
- Microservices: `Kratos`, `Kitex`, `go-zero`
- High-performance networking: `gnet`
- ORM: `Ent`
- Testing: `Ginkgo`
