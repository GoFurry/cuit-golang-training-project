package main

import (
	"fmt"
	"math"
	"reflect"
)

/**
 * title: Bool demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// Go语言不存在隐式类型转换, 所有的类型转换都必须显式声明

func main() {
	// 取值范围较小的类型转换到一个取值范围较大的类型
	var i16 int16 = 0x7068
	fmt.Println(reflect.TypeOf(i16), reflect.TypeOf(int8(i16))) // int16 int8
	fmt.Println(i16, int8(i16))                                 // 精度丢失 28776 104

	// 只有相同底层类型的变量之间可以进行相互转换
	//bool(i16) // 编译报错 cannot convert i16 (variable of type int16) to type bool
	fmt.Println(string(i16)) // 灨

	// 浮点数转整数只会保留整数部分
	pi := math.Pi
	fmt.Println(int(pi)) // 3
}
