**Read this in other languages:**
[English](README_en.md) | [中文](README.md)

# 02-go-core: Go Core Features Module

## Module Description

This module focuses on Go core features and commonly used advanced syntax, including arrays, slices, maps, structs, methods, interfaces, embedding, error handling, defer/panic/recover, concurrency, and synchronization. Each demo includes clear comments for step-by-step learning after finishing the basics.

## Learning Path (Recommended Order)

001-demo-Array (Array)

002-demo-Slice (Slice)

003-demo-Map (Map)

004-demo-Struct (Struct)

005-demo-Method (Method)

006-demo-Interface (Interface)

007-demo-Embedding (Embedding)

008-demo-ErrorHandling (Error Handling)

009-demo-DeferPanicRecover (Defer/Panic/Recover)

010-demo-Goroutine (Goroutine)

011-demo-Channel (Channel)

012-demo-Sync (Concurrency Sync)

## Core Content of Each Demo

| No.  | Directory Name     | Core Knowledge Points                                           | Learning Focus                                      |
| ---- | ------------------ | -------------------------------------------------------------- | --------------------------------------------------- |
| 001  | Array              | Array declaration, length as part of type, traversal & copy     | Difference between array and slice, pointer passing |
| 002  | Slice              | Slice header, len/cap, append/copy, shared underlying array     | Growth behavior and shared data effects             |
| 003  | Map                | Map declaration/init, read/write/delete, ok check, iteration    | nil map vs empty map, unordered iteration           |
| 004  | Struct             | Struct declaration, zero value, nested structs, pointers        | Organizing data with composition                    |
| 005  | Method             | Method declaration, value vs pointer receiver, method set       | When to use pointer receivers                       |
| 006  | Interface          | Interface definition, implicit implementation, type assertions  | Polymorphism and safe assertions                    |
| 007  | Embedding          | Struct embedding, promoted fields, conflict resolution          | Composition over inheritance                         |
| 008  | ErrorHandling      | error interface, errors.Is/As, wrapping, custom errors          | Error chain matching and semantics                  |
| 009  | DeferPanicRecover  | defer order, panic/recover, resource cleanup                    | LIFO behavior and recovery                          |
| 010  | Goroutine          | Creating goroutines, concurrent execution, waiting             | Concurrency mindset and coordination                |
| 011  | Channel            | Unbuffered/buffered channels, close, range receive              | Blocking rules and communication patterns           |
| 012  | Sync               | WaitGroup, Mutex, safe counters                                 | Protecting shared data and synchronization          |

## Execution & Debugging

### 1. Run a Single Demo

Navigate to the target demo directory and run:

```
# Example: run Array
cd 001-demo-Array 数组
go run array.go
```

### 2. Common Issue Resolution

- **"imported and not used"**: Remove unused imports or actually use the package.

- **"all goroutines are asleep - deadlock!"**: Channel send/receive mismatch, check blocking pairs.

- **Inconsistent concurrent results**: Protect shared variables with Mutex or other sync tools.

## Learning Recommendations

**Understand, then modify**: Read the comments first, then tweak slice capacity, channel buffer size, or concurrency count to observe changes.

**Focus on errors & concurrency**: These are core Go design points and worth repeated practice.

## Supporting Resources

- Not available for now
