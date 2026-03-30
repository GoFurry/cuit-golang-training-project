package main

import "fmt"

/**
 * title: DeferPanicRecover demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// defer 会在函数返回前执行，遵循后进先出（LIFO）
// panic 用于异常中断，recover 可以在 defer 中恢复

func main() {
	defer fmt.Println("defer #1")
	defer fmt.Println("defer #2")
	fmt.Println("main start")

	res, err := safeDivide(10, 0)
	fmt.Println("result:", res, "err:", err)

	fmt.Println("main end")
}

// safeDivide 使用 defer + recover 处理 panic
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover from panic: %v", r)
		}
	}()

	if b == 0 {
		panic("divide by zero")
	}
	return a / b, nil
}
