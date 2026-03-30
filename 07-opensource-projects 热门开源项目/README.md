**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 07-opensource-projects：热门开源项目

## 模块说明

这一章不再放单点 Demo，而是整理一份适合 Go 学习者长期追踪的开源项目清单。

截至 `2026-03-30`，下面这份列表优先筛选：

- 在 Go 社区影响力长期稳定的项目
- 有真实工程价值、不是一次性玩具项目
- 源码结构、文档质量或架构思路值得学习
- 能帮助你从“会写语法”走向“会看工程”

这不是绝对的 Star 排名，也不是唯一答案，而是一份更偏“学习路线图”的推荐清单。

## 推荐阅读顺序

- 入门阶段：`Gin`、`Chi`、`Cobra`、`Hugo`
- 工程进阶：`Caddy`、`Traefik`、`Prometheus`、`MinIO`
- 架构进阶：`Kubernetes`、`etcd`、`containerd`、`NATS`
- 前沿拓展：`Milvus`、`Loki`、`Ollama`、`go-ethereum`

## 热门开源项目清单

| 类别 | 项目 | 地址 | 短评 | 推荐原因 |
| ---- | ---- | ---- | ---- | ---- |
| Web 框架 | Gin | [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) | Go Web 框架里的高频选手，API 直观，上手极快。 | 适合学习路由、中间件、请求绑定、JSON API 组织方式，也是很多 Go Web 项目的起点。 |
| Web 路由 | Chi | [github.com/go-chi/chi](https://github.com/go-chi/chi) | 很“标准库味”的轻量路由器。 | 适合理解更贴近 `net/http` 的 Go 风格写法，能帮助你建立比“只会框架”更稳的 Web 基础。 |
| CLI 工具 | Cobra | [github.com/spf13/cobra](https://github.com/spf13/cobra) | Go CLI 生态里的事实标准之一。 | 很多知名工具都在用它，读它能学到子命令组织、参数系统、帮助文档生成和工程化 CLI 设计。 |
| 静态站点 | Hugo | [github.com/gohugoio/hugo](https://github.com/gohugoio/hugo) | 成熟度非常高的静态站点生成器。 | 适合学习大型 CLI 程序、文件系统处理、模板系统和性能优化，是“Go 做工具”的代表作。 |
| Web Server | Caddy | [github.com/caddyserver/caddy](https://github.com/caddyserver/caddy) | 配置体验很现代的 Web 服务器，自动 HTTPS 很出名。 | 适合学习模块化架构、配置驱动设计、插件系统以及实际可落地的服务端工程写法。 |
| 云原生 | Kubernetes | [github.com/kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) | 云原生世界里最有代表性的 Go 项目之一。 | 适合在进阶阶段学习大型仓库组织、API Machinery、控制器模式、声明式系统设计。 |
| 分布式存储 | etcd | [github.com/etcd-io/etcd](https://github.com/etcd-io/etcd) | 小而硬核的分布式键值存储。 | 非常适合拿来理解 Raft、一致性、租约、Watch 机制，以及 Go 如何写分布式基础组件。 |
| 可观测性 | Prometheus | [github.com/prometheus/prometheus](https://github.com/prometheus/prometheus) | 监控领域的核心项目，生态影响力极强。 | 适合学习指标采集、规则计算、服务发现、时序数据处理和长期演进的工程结构。 |
| 网关 / 反向代理 | Traefik | [github.com/traefik/traefik](https://github.com/traefik/traefik) | 云原生场景下很常见的网关和反向代理。 | 适合学习动态配置、Provider 抽象、中间件链路和网络流量入口层的工程实践。 |
| 容器运行时 | containerd | [github.com/containerd/containerd](https://github.com/containerd/containerd) | 容器生态里的底层核心组件。 | 适合理解容器运行时、守护进程设计、任务生命周期和云原生底座的真实实现。 |
| 消息系统 | NATS Server | [github.com/nats-io/nats-server](https://github.com/nats-io/nats-server) | 轻快、干净、性能很强的消息系统。 | 源码相对清爽，适合学习协议实现、并发网络编程、发布订阅模型和高性能服务设计。 |
| 对象存储 | MinIO | [github.com/minio/minio](https://github.com/minio/minio) | S3 兼容对象存储领域的明星项目。 | 适合学习 API 兼容设计、存储服务架构、多磁盘/分布式场景以及非常实战的工程代码。 |
| 向量数据库 | Milvus | [github.com/milvus-io/milvus](https://github.com/milvus-io/milvus) | AI 时代很有代表性的向量数据库项目。 | 如果你想把 Go 学习和 AI 基础设施结合起来，它能帮助你理解大型数据系统与检索系统的实现。 |
| 日志系统 | Loki | [github.com/grafana/loki](https://github.com/grafana/loki) | Grafana 生态里的日志聚合核心项目。 | 适合学习日志索引策略、分布式组件拆分、可观测性工程和与 Prometheus 风格一致的产品设计。 |
| 大模型工具 | Ollama | [github.com/ollama/ollama](https://github.com/ollama/ollama) | 本地运行大模型的热门工具，工程体验很强。 | 适合了解 Go 在 AI 工具链中的角色，包括模型管理、本地服务封装、API 暴露与跨平台分发。 |
| 区块链 | go-ethereum | [github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) | 区块链世界里最知名的 Go 项目之一。 | 适合挑战大型复杂系统，学习 P2P 网络、交易池、区块同步、虚拟机和长生命周期项目维护方式。 |

## 怎么读这些项目

- 不要一上来就“从第一行读到最后一行”，先读 README、目录结构、启动入口。
- 先看 `main`、配置加载、路由注册、模块初始化，再往里追具体实现。
- 优先挑和你当前学习目标一致的项目，比如学 Web 就先看 Gin / Chi，学分布式就先看 etcd / NATS。
- 看源码时记笔记，尤其记录“项目是怎么分层的”“公共抽象放在哪”“配置和运行时状态怎么流动”。

## 推荐关注点

### 1. 学 Go Web 工程

- Gin
- Chi
- Caddy
- Traefik

### 2. 学 CLI 与工具开发

- Cobra
- Hugo
- Ollama

### 3. 学云原生与分布式系统

- Kubernetes
- etcd
- containerd
- NATS
- Prometheus

### 4. 学存储、检索与数据基础设施

- MinIO
- Milvus
- Loki

## 额外建议

- 如果你是第一次看大型开源项目，建议先选 `Cobra`、`Gin`、`Chi`、`Hugo` 这类入口更友好的仓库。
- 如果你已经完成本仓库前面的基础语法、核心特性、标准库和框架部分，再去读 `Kubernetes`、`etcd`、`go-ethereum` 会更顺。
- 真正的成长不在于“看过多少仓库”，而在于你能不能总结出它们在目录设计、接口抽象、并发模型、错误处理上的共性。
