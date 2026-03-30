package main

import "fmt"

/**
 * title: Array demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 数组长度是类型的一部分，[3]int 和 [4]int 是不同类型
// 数组是值类型，赋值和传参会发生拷贝
// len(array) 获取长度，cap(array) 对数组等于 len(array)

func main() {
	// 声明并使用零值数组
	var a [3]int
	fmt.Println("a =", a, "len =", len(a), "cap =", cap(a))

	// 字面量声明
	b := [3]int{1, 2, 3}
	c := [...]int{4, 5, 6, 7} // ... 让编译器推断长度
	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	// 修改元素
	b[0] = 100
	fmt.Println("b =", b)

	// 遍历数组
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]=%d ", i, c[i])
	}
	fmt.Println()

	for idx, val := range c {
		fmt.Printf("idx=%d val=%d ", idx, val)
	}
	fmt.Println()

	// 传参拷贝演示
	modifyArrayCopy(b)
	fmt.Println("after modifyArrayCopy b =", b)

	// 传指针可以修改原数组
	modifyArrayByPointer(&b)
	fmt.Println("after modifyArrayByPointer b =", b)
}

// modifyArrayCopy 传值会发生拷贝
func modifyArrayCopy(arr [3]int) {
	arr[0] = -1
	fmt.Println("modifyArrayCopy arr =", arr)
}

// modifyArrayByPointer 传指针修改原数组
func modifyArrayByPointer(arr *[3]int) {
	(*arr)[0] = -2
	fmt.Println("modifyArrayByPointer arr =", *arr)
}
