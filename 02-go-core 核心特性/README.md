### 中文 [README.md](README.md)

### English [README_en.md](README_en.md)

# 02-go-core：Go 语言核心特性模块

## 模块说明

本目录聚焦 Go 语言的核心特性与常用高级语法，覆盖数组、切片、映射、结构体、方法、接口、组合嵌入、错误处理、延迟与异常、并发与同步等关键知识点。所有 Demo 均配有注释说明，适合在掌握基础语法后进阶学习。

## 学习路径（推荐顺序）

001-demo-Array（数组）

002-demo-Slice（切片）

003-demo-Map（映射）

004-demo-Struct（结构体）

005-demo-Method（方法）

006-demo-Interface（接口）

007-demo-Embedding（组合嵌入）

008-demo-ErrorHandling（错误处理）

009-demo-DeferPanicRecover（延迟与异常）

010-demo-Goroutine（协程）

011-demo-Channel（通道）

012-demo-Sync（并发同步）

## 各 Demo 核心内容

| 序号 | 目录名称           | 核心知识点                                                | 学习重点                                         |
| ---- | ------------------ | --------------------------------------------------------- | ------------------------------------------------ |
| 001  | Array              | 数组声明、长度是类型一部分、遍历与传参拷贝                | 数组与切片的区别、指针传参修改原数组            |
| 002  | Slice              | 切片结构、len/cap、append/copy、共享底层数组             | 扩容规则与共享数据的影响                        |
| 003  | Map                | map 声明/初始化、读写删除、ok 判断、遍历                 | nil map 与空 map 区别、遍历无序性               |
| 004  | Struct             | 结构体声明、零值、嵌套结构体、指针与匿名结构体           | 组合结构体数据组织方式                          |
| 005  | Method             | 方法声明、值接收者/指针接收者、方法集                     | 何时使用指针接收者                               |
| 006  | Interface          | 接口定义、隐式实现、空接口、类型断言/类型选择            | 多态使用与断言安全性                             |
| 007  | Embedding          | 结构体嵌入、字段提升、字段冲突处理                        | 组合优于继承的设计思路                           |
| 008  | ErrorHandling      | error 接口、errors.Is/As、错误包装与自定义错误           | 错误链判断与错误语义表达                         |
| 009  | DeferPanicRecover  | defer 执行顺序、panic/recover、资源释放                   | LIFO 与恢复机制                                  |
| 010  | Goroutine          | goroutine 创建、并发执行、等待退出                        | 并发思维与主协程同步                             |
| 011  | Channel            | 无缓冲/有缓冲通道、关闭通道、range 读取                   | 通道阻塞规则与通信模式                           |
| 012  | Sync               | WaitGroup、Mutex、并发安全计数                           | 共享数据保护与并发同步                           |

## 运行与调试

### 1. 单个 Demo 运行

进入目标 Demo 目录，执行 go run 命令：

```
# 示例：运行 Array
cd 001-demo-Array 数组
go run array.go
```

### 2. 常见问题解决

- **"imported and not used"**：导入了未使用的包，删除无用 import 或使用该包。

- **"all goroutines are asleep - deadlock!"**：通道收发不匹配导致阻塞，检查发送/接收是否成对。

- **并发结果不一致**：共享变量未加锁，使用 Mutex 或其他同步手段保护。

## 学习建议

**先理解概念再动手**：读完注释后运行，再尝试修改切片容量、通道缓冲区或并发数量观察差异。

**重视错误与并发**：错误处理和并发是 Go 的核心设计，建议反复练习。

## 配套资源

- 暂无
