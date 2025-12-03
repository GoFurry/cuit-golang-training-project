package main

import "fmt"

/**
 * title: Function demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

func main() {
	// 匿名函数
	foo := func() {
		fmt.Println("foo")
	}
	foo() // 调用匿名函数 输出 foo

	// 在末尾可以赋值入参
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(sum) // 输出 3

	// 闭包
	accumulator := acc(1)
	accumulator() // acc 2
	accumulator() // acc 3
}

// f1 与 f2 两个函数的入参是相同的
func f1(i, j, k int, s, t string) {

}
func f2(i int, j int, k int, s string, t string) {

}

func add(x int, y int) int       { return x + y }
func sub(x, y int) (z int)       { z = x - y; return } // 入参和返回值都可以一开始就声明变量
func first(x int, _ int) int     { return x }          // 可以用匿名函数
func zero(int, int) int          { return 0 }          // 入参和返回值可只用类型
func typedTwoValues() (int, int) { return 1, 2 }       // 可以有多个返回值

func acc(num int) func() int {
	return func() int {
		num++
		fmt.Println("acc", num)
		return num
	}
}
