package main

import "fmt"

/**
 * title: Const demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// 常量用于存储不会改变的数据

func main() {
	const cat string = "cat"
	const dog = "dog"
	const num = 1 + 2
	// 常量的值必须在编译期间就能获得
	//const number = GetNumber() // 报错 GetNumber() (value of type int) is not constant

	// 批量定义常量 第一个变量必须赋值
	const (
		x = 10
		y
		z = 20
		g
	)
	fmt.Println(y, g) // 10 20

	// iota 枚举类型
	type Weekday int // 别名 Weekday = int
	const (
		Sunday    Weekday = iota // 0
		Monday                   // 1
		Tuesday                  // 2
		Wednesday                // 3
		Thursday                 // 4
		Friday                   // 5
		Saturday                 // 6
	)
	fmt.Printf("%T\n", Sunday) // main.Weekday

	// 常量可以构成类型的一部分
	var numx [x]int
	var numy [10]int
	numx, numy = numx, numy

}

func GetNumber() int {
	return 1
}
