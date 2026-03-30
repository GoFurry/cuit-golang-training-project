**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 04-go-frameworks：Go 主流框架应用模块
## 模块说明
本目录涵盖 Go 生态中的主流框架，从 Web 开发到微服务、从命令行工具到桌面应用，提供完整的框架学习路径。每个框架都配备独立的可运行 Demo，展示核心特性与实际应用场景。
## 主题与目录结构
```
04-go-frameworks/
├── 01-web-frameworks Web框架/
├── 02-microservices 微服务框架/
├── 03-high-performance-http 高性能HTTP框架/
├── 04-cli-frameworks CLI框架/
├── 05-gui-frameworks GUI框架/
├── 06-crawler-frameworks 爬虫框架/
├── 07-orm-frameworks ORM框架/
└── 08-testing-frameworks 测试框架/
```
## 各主题 Demo 清单
### 01-web-frameworks Web框架
- Gin（最流行的轻量级Web框架）
- Echo（高性能Web框架）
- Fiber（Express.js风格的Web框架）
- Chi（轻量级路由器）
### 02-microservices 微服务框架
- gRPC（Google高性能RPC框架）
- Kratos（B站开源微服务框架）
- Kitex（字节跳动微服务RPC框架）
- go-zero（微服务框架）
### 03-high-performance-http 高性能HTTP框架
- gnet（事件驱动网络框架）
- fasthttp（高性能HTTP实现）
### 04-cli-frameworks CLI框架
- Cobra（强大的CLI应用构建库）
### 05-gui-frameworks GUI框架
- Wails（现代桌面应用开发框架）
### 06-crawler-frameworks 爬虫框架
- Colly（优雅的网络爬虫框架）
### 07-orm-frameworks ORM框架
- GORM（功能完整的ORM库）
- XORM（简单强大的ORM引擎）
### 08-testing-frameworks 测试框架
- Testify（丰富的测试工具集）
- GoConvey（BDD风格测试框架）
## 运行方式
### 1. 运行单个 Demo
```
# 示例：运行 Gin demo
cd 01-web-frameworks Web框架/001-demo-gin
go mod tidy
go run main.go
```
### 2. 微服务类 Demo
```
# gRPC 需要先生成代码
cd 02-microservices 微服务框架/001-demo-grpc
go generate
go run server/main.go
# 另开终端运行客户端
go run client/main.go
```
### 3. GUI 应用 Demo
```
cd 05-gui-frameworks GUI框架/001-demo-wails
wails build
# 或开发模式
wails dev
```
## 学习建议
- Web框架建议从 Gin 开始，掌握基础概念后再学习其他框架。
- 微服务框架需要先理解 RPC 概念，建议先学习 gRPC。
- 高性能框架适合有一定基础后学习，重点理解性能优化思路。
- GUI 和爬虫框架可根据实际需求选择性学习。
- ORM 框架建议结合数据库实践，先掌握 SQL 基础。
## 配套资源
- 各框架官方文档链接
- 性能对比测试结果
- 最佳实践指南