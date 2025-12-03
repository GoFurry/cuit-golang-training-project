package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
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

	// 字符串转数值类型
	intNumber1, _ := strconv.Atoi("8848")
	intNumber2, _ := strconv.ParseInt("4399", 10, 64)
	intNumber3, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(intNumber1, intNumber2, intNumber3)                                                 // 8848 4399 3.1415926
	fmt.Println(reflect.TypeOf(intNumber1), reflect.TypeOf(intNumber2), reflect.TypeOf(intNumber3)) // int int64 float64

	// 数值类型转字符串
	strNumber1 := strconv.FormatInt(007, 10)
	strNumber2 := fmt.Sprintf("%d", 996)
	fmt.Println(strNumber1, strNumber2)                                 // 007 996
	fmt.Println(reflect.TypeOf(strNumber1), reflect.TypeOf(strNumber2)) // string string
}
