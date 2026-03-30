package main

import "fmt"

/**
 * title: Embedding demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 嵌入（组合）可以将字段和方法“提升”到外层结构体

type Base struct {
	ID int
}

func (b Base) PrintID() {
	fmt.Println("ID =", b.ID)
}

type Order struct {
	Base
	Amount int
}

// 字段名冲突示例

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

func main() {
	order := Order{Base: Base{ID: 1001}, Amount: 99}
	order.PrintID() // 直接访问 Base 的方法
	fmt.Println("order ID =", order.ID)
	fmt.Println("amount =", order.Amount)

	c := C{A: A{Name: "Alice"}, B: B{Name: "Bob"}}
	// c.Name 会产生歧义，需要显式指定
	fmt.Println("c.A.Name =", c.A.Name)
	fmt.Println("c.B.Name =", c.B.Name)
}
