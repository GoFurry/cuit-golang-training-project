package main

import (
	"fmt"
	"math"
)

/**
 * title: Number types demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// 数值类型分为: 整数、浮点数、复数

// 整型包括: int8/int16/int32/int64 uint8/uint16/uint32/uint64 rune/byte int/uint uintptr
// u 开头的是无符号整形
// rune 和 int32 等价, byte 和 int8 等价
// int/uint 由于编译器和计算机硬件的不同 实际大小在 32-bit 和 64-bit 之间变化
// uintptr 无固定大小但是可以容纳指针, 用于底层编程
// n-bit 的 int 取值范围 -2^(n-1) 到 2^(n-1)-1
// n-bit 的 uint 取值范围 0 到 2^n-1

// 浮点数: float32/float64
// float32 最大值约 3.4e38 最小值 1.4e-45
// float64 最大值约 1.8e308 最小值 4.9e-324

// 复数类型很少使用可暂时跳过
// 复数类型: complex128(64 位实数和虚数) / complex64(32 位实数和虚数)
// complex128 为复数的默认类型

func main() {
	// 整形极值
	var (
		u8  uint8
		u16 uint16
		u32 uint32
		u64 uint64
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)
	u8, u16, u32, u64 = 0xff, 0xffff, 0xffffffff, 0xffffffffffffffff
	fmt.Printf("u8=%d u16=%d u32=%d u64=%d\n", u8, u16, u32, u64)
	i8, i16, i32, i64 = 0x7f, 0x7fff, 0x7fffffff, 0x7fffffffffffffff
	fmt.Printf("i8=%d i16=%d i32=%d i64=%d\n", i8, i16, i32, i64)
	i8, i16, i32, i64 = -0x7f-1, -0x7fff-1, -0x7fffffff-1, -0x7fffffffffffffff-1
	fmt.Printf("i8=%d i16=%d i32=%d i64=%d\n", i8, i16, i32, i64)

	// 浮点数极值
	fmt.Println(math.MaxFloat64, math.MaxFloat32)
	// 浮点数精度问题
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // 输出 "true"
	// 可以只写小数或整数部分
	e := .71828
	h := 1.
	fmt.Println(e, h) // 0.71828 1
	// 科学计数法
	Avogadro := 6.02214129e23     // 阿伏伽德罗常数
	Planck := 6.62606957e-34      // 普朗克常数
	fmt.Println(Avogadro, Planck) // 0.71828 1
	// 格式化输出
	fmt.Printf("%f\n", math.Pi)   // 3.141593
	fmt.Printf("%.2f\n", math.Pi) // 3.14

	// 复数声明
	z := complex(1, 2)
	fmt.Println(z)       // (1+2i)
	fmt.Println(real(z)) // 实部 1
	fmt.Println(imag(z)) // 虚部 2
}
