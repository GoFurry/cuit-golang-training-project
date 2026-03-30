**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 04-go-frameworks：Go 主流框架

## 模块说明

这一章聚焦 Go 生态里最常见、最值得入门的框架与工具链，目标不是堆砌大而全的工程，而是提供一组可以直接运行、便于对比、适合教学的最小 Demo。

本章所有带三方依赖的示例都采用“每个 Demo 一个独立模块”的方式组织，并在 `go.mod` 中锁定：

```go
go 1.26
toolchain go1.26.0
```

这样既方便单独学习，也不会让不同框架之间的依赖相互干扰。

## 目录结构

```text
04-go-frameworks 主流框架/
├── 01-web-frameworks Web框架/
│   ├── 001-demo-gin
│   ├── 002-demo-echo
│   ├── 003-demo-fiber
│   └── 004-demo-chi
├── 02-microservices 微服务框架/
│   └── 001-demo-grpc
├── 03-high-performance-http 高性能HTTP框架/
│   └── 001-demo-fasthttp
├── 04-cli-frameworks CLI框架/
│   └── 001-demo-cobra
├── 05-gui-frameworks GUI框架/
│   └── 001-demo-wails
├── 06-crawler-frameworks 爬虫框架/
│   └── 001-demo-colly
├── 07-orm-frameworks ORM框架/
│   ├── 001-demo-gorm
│   └── 002-demo-xorm
└── 08-testing-frameworks 测试框架/
    ├── 001-demo-testify
    └── 002-demo-goconvey
```

## Demo 清单

### 01-web-frameworks Web框架

- `001-demo-gin`：学习路由、中间件、JSON 绑定与响应
- `002-demo-echo`：学习 Context 封装、参数解析和错误处理
- `003-demo-fiber`：学习 Express 风格 API 与高性能 HTTP 开发体验
- `004-demo-chi`：学习轻量路由、标准库风格中间件

### 02-microservices 微服务框架

- `001-demo-grpc`：使用 `proto3` 定义服务，生成客户端和服务端代码

### 03-high-performance-http 高性能HTTP框架

- `001-demo-fasthttp`：体验比 `net/http` 更底层的高性能 HTTP 处理方式

### 04-cli-frameworks CLI框架

- `001-demo-cobra`：构建多命令命令行工具，体验子命令、参数解析与帮助信息

### 05-gui-frameworks GUI框架

- `001-demo-wails`：使用 Go + 前端技术栈构建桌面应用

### 06-crawler-frameworks 爬虫框架

- `001-demo-colly`：使用本地测试站点演示 HTML 抓取、链接跟进和数据提取

### 07-orm-frameworks ORM框架

- `001-demo-gorm`：学习模型定义、迁移、插入、查询、更新
- `002-demo-xorm`：对比另一套 ORM 风格，理解会话与引擎的基本使用

### 08-testing-frameworks 测试框架

- `001-demo-testify`：学习断言、错误校验和测试组织方式
- `002-demo-goconvey`：体验 BDD 风格测试表达

## 运行方式

### 1. 运行 Web / HTTP / CLI / ORM / 爬虫 Demo

```bash
cd "01-web-frameworks Web框架/001-demo-gin"
go mod tidy
go run .
```

### 2. 运行 gRPC Demo

先生成代码，再分别运行服务端和客户端：

```bash
cd "02-microservices 微服务框架/001-demo-grpc"
go generate ./...
go mod tidy
go run ./server
```

打开另一个终端：

```bash
cd "02-microservices 微服务框架/001-demo-grpc"
go run ./client
```

### 3. 运行测试框架 Demo

```bash
cd "08-testing-frameworks 测试框架/001-demo-testify"
go mod tidy
go test ./...
```

### 4. 运行 Wails Demo

```bash
cd "05-gui-frameworks GUI框架/001-demo-wails"
go mod tidy
wails dev
```

## 学习建议

- Web 框架建议先学 `Gin`，再用 `Echo / Fiber / Chi` 对比不同的 API 设计风格。
- 微服务建议先掌握 `gRPC` 的接口定义、代码生成和服务注册，再去学 Kratos、Kitex、go-zero 这类工程框架。
- ORM 建议结合 SQLite 或 MySQL 一起练习，重点体会“模型定义”和“SQL 抽象”的边界。
- 测试框架建议把断言库与标准库 `testing` 结合起来理解，而不是只记住语法。
- Wails 更适合作为“Go 调用前端”的入门 GUI 方案，等你熟悉以后再扩展到完整桌面项目。

## 后续扩展建议

这一章先优先放“轻量、可运行、便于理解”的 Demo。等你完成本章后，下一轮可以继续扩展：

- Web：`Beego`、`Hertz`
- 微服务：`Kratos`、`Kitex`、`go-zero`
- 高性能网络：`gnet`
- ORM：`Ent`
- 测试：`Ginkgo`
