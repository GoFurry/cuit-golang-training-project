### 中文 [README.md](README.md)

### English [README_en.md](README_en.md)

# 03-go-standard-lib：Go 标准库实践模块

## 模块说明

本目录以主题形式组织 Go 标准库的常用包，强调“可运行的小 Demo”。每个包对应一个独立 Demo，适合动手验证标准库能力与使用场景。

## 主题与目录结构

```
03-go-standard-lib/
├── 01-io-filesystem I-O与文件系统/
├── 02-encoding-parsing 编码与解析/
├── 03-network-protocol 网络与协议/
├── 04-concurrency-sync 并发与同步/
├── 05-errors-logging 错误与日志/
├── 06-strings-text 字符串与文本/
├── 07-ds-algorithms 数据结构与算法/
├── 08-system-runtime 系统与运行时/
├── 09-security-crypto 安全与加密/
└── 10-database 数据库/
```

## 各主题 Demo 清单

### 01-io-filesystem I-O与文件系统

- os
- io
- bufio
- path/filepath
- embed

### 02-encoding-parsing 编码与解析

- encoding/json
- encoding/csv
- encoding/base64
- encoding/hex
- encoding/xml
- text/template
- html/template

### 03-network-protocol 网络与协议

- net/http
- net/url
- net
- mime/multipart
- net/http/httptest

### 04-concurrency-sync 并发与同步

- context
- sync
- sync/atomic
- time

### 05-errors-logging 错误与日志

- errors
- log
- log/slog

### 06-strings-text 字符串与文本

- strings
- bytes
- regexp
- strconv
- unicode/utf8

### 07-ds-algorithms 数据结构与算法

- sort
- container/heap
- container/list
- container/ring

### 08-system-runtime 系统与运行时

- flag
- runtime
- testing（使用 go test 运行）
- reflect

### 09-security-crypto 安全与加密

- crypto/rand
- crypto/sha256
- crypto/hmac
- crypto/tls

### 10-database 数据库

- database/sql（需配合具体驱动）

## 运行方式

### 1. 运行单个 Demo

```
# 示例：运行 json demo
cd 02-encoding-parsing 编码与解析/001-demo-json
go run json.go
```

### 2. 测试类 Demo（testing）

```
cd 08-system-runtime 系统与运行时/003-demo-testing
go test -v
```

## 学习建议

- 每次只关注一个包：先读注释，再改代码观察输出变化。
- 网络/并发类 Demo 建议多跑几次，理解阻塞与时序。
- database/sql 需引入数据库驱动后再深入练习。

## 配套资源

- 暂无
