package main

import "fmt"

/**
 * title: Pointer demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// 指针分为两个核心概念
// 类型指针: 允许对这个指针类型的数据进行修改, 类型指针不能进行偏移和运算。
// 切片: 由指向起始元素的原始指针、元素数量和容量组成

// 指针在 32位/64位 的计算机上分别占 4个/8个 字节
// 默认值为 nil

func main() {
	// & 用于取变量的地址
	cat, dog := 666, "dog"
	fmt.Println(&cat, &dog) // cat 的地址 dog 的地址

	// &ptr 用于取指针指向的地址
	// *ptr 用于取指针指向的值
	ptr := &cat
	fmt.Printf("%p %T\n", ptr, ptr) // ptr的地址(和cat地址相同) ptr的类型
	fmt.Println(*ptr)               // 输出 666

	// 函数传入指针地址可抑制就修改变量的值
	x, y := 10, 20
	fmt.Println(x, y) // 10 20
	swap(x, y)
	fmt.Println(x, y) // 10 20
	ptrSwap(&x, &y)
	fmt.Println(x, y) // 20 10

	// new 创建指针
	str := new(string)
	*str = dog
	fmt.Println(*str) // 输出 dog

}

// ptrSwap 指针形式交换ab值
func ptrSwap(a, b *int) {
	*a, *b = *b, *a
}

// swap 交换ab值
func swap(a, b int) {
	a, b = b, a
}
