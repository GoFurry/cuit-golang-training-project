package main

import "fmt"

/**
 * title: Slice demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 切片是对数组的一个动态视图，包含指针、长度、容量
// nil 切片与空切片不同：nil 切片的底层数组为空

func main() {
	// nil 切片与空切片
	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("nilSlice == nil:", nilSlice == nil)
	fmt.Println("emptySlice == nil:", emptySlice == nil)

	// 字面量与 make 创建
	literals := []int{1, 2, 3, 4}
	made := make([]int, 3, 5) // len=3 cap=5，默认值为 0
	fmt.Println("literals =", literals, "len=", len(literals), "cap=", cap(literals))
	fmt.Println("made =", made, "len=", len(made), "cap=", cap(made))

	// append 追加元素，容量不够时会扩容并重新分配底层数组
	made = append(made, 10, 11)
	fmt.Println("made after append =", made, "len=", len(made), "cap=", cap(made))

	// 切片共享底层数组
	s1 := literals[1:3] // [2,3]
	s1[0] = 200
	fmt.Println("s1 =", s1)
	fmt.Println("literals after s1 modify =", literals)

	// copy 复制数据到目标切片
	dst := make([]int, 2)
	copy(dst, literals)
	fmt.Println("dst =", dst)

	// 重新切片
	s2 := literals[:2]
	s3 := literals[2:]
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)
}
