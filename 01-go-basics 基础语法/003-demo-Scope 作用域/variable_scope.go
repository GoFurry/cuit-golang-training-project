package main

import "fmt"

/**
 * title: variable-scope demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// 变量的作用域分为三种:
//   函数内定义的变量称为 局部变量
//   函数外定义的变量称为 全局变量
//   函数定义中的变量称为 形式参数

// 堆heap: 堆是用于存放进程执行中被动态分配的内存段
// 栈stack: 栈用来存放程序暂时创建的局部变量
// 实际开发中并不需要刻意的实现变量的逃逸行为 (栈上的变量逃逸到堆上)
// 但逃逸的变量需要额外分配内存, 对性能优化可能会产生细微的影响
// 逃逸分析 go run -gcflags=-m ./variable_scope.go

var globalA = 10

func main() {
	// 局部变量
	a, b := 1, 2
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c) // 输出 a = 1, b = 2, c = 3

	// 局部变量会覆盖全局变量
	fmt.Println(globalA) // 输出 10
	globalA := 3.14      // 局部变量
	fmt.Println(globalA) // 输出 3.14
	demo(a, b)
	fmt.Println(globalA, a, b) // 输出 3.14 1 2

}

// 函数传入的参数叫做形式参数
func demo(x int, y int) {
	fmt.Println(globalA) // 读取到的是全局变量 输出 10
	globalA = x          // 全局变量修改不影响局部变量
	x++
	y++
}
