**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 03-go-standard-lib: Go Standard Library Practice Module

## Module Description

This module organizes frequently used Go standard library packages by themes. Each package has a small runnable demo to help you learn by doing.

## Themes & Structure

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

## Demo List by Theme

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
- testing (run with go test)
- reflect

### 09-security-crypto 安全与加密

- crypto/rand
- crypto/sha256
- crypto/hmac
- crypto/tls

### 10-database 数据库

- database/sql (requires a driver)

## How To Run

### 1. Run a single demo

```
# Example: run json demo
cd 02-encoding-parsing 编码与解析/001-demo-json
go run json.go
```

### 2. Testing demo

```
cd 08-system-runtime 系统与运行时/003-demo-testing
go test -v
```

## Learning Tips

- Focus on one package at a time: read comments first, then tweak and rerun.
- For network/concurrency demos, run multiple times to observe timing and blocking.
- For database/sql, add a driver to practice deeper.

## Supporting Resources

- Not available for now
