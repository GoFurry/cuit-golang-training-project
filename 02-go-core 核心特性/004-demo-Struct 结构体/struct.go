package main

import "fmt"

/**
 * title: Struct demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 结构体用于组织不同类型的数据
// 结构体零值字段为对应类型零值

type Person struct {
	Name string
	Age  int
}

type Address struct {
	City  string
	Street string
}

type Employee struct {
	Person  Person
	Address Address
	Level   int
}

func main() {
	// 零值
	var p1 Person
	fmt.Println("p1 =", p1)

	// 字面量初始化
	p2 := Person{Name: "Alice", Age: 20}
	fmt.Println("p2 =", p2)

	// 指针方式初始化
	p3 := &Person{Name: "Bob", Age: 22}
	fmt.Println("p3 =", p3)

	// 匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "Cindy", Age: 18}
	fmt.Println("p4 =", p4)

	// 组合结构体
	e := Employee{
		Person:  Person{Name: "David", Age: 25},
		Address: Address{City: "Chengdu", Street: "Renmin Road"},
		Level:   3,
	}
	fmt.Println("employee =", e)
	fmt.Println("employee city =", e.Address.City)
}
