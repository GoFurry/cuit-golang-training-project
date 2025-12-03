package main

import "fmt"

/**
 * title: variable demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// 变量的声明如 var name1, name2 type
// 其中 var 是声明变量的关键字， name 是变量名, type 是变量的类型
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // uint8 的别名
// rune // int32 的别名 代表一个 Unicode 码
// float32 float64
// complex64 complex128

// 当变量被声明后, 系统自动赋予它该类型的零值
// int 为 0, float 为 0.0, bool 为 false, string 为空字符串, 指针为 nil 等
// 变量的命名规则遵循骆驼命名法, 例如: stuNumber 和 stuAge

func main() {
	// 声明单个变量
	var num int
	fmt.Println(num) // 输出 0

	// 赋值声明
	var sum = 5
	fmt.Println(sum) // 输出 5

	// 简短声明 名字 := 表达式
	// 编译器会自动根据右值类型推断出左值的对应类型
	str := "abc"
	fmt.Println(str) // 输出 abc
	var1, var2 := 10, 15
	fmt.Println(var1, var2) // 输出 10 15
	var1, var2 = var2, var1
	fmt.Println(var1, var2) // 输出 15 10

	// 批量声明
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	fmt.Printf("a=%d b=%s c=%f d=%v e=%v\n", a, b, c, d, e) // 输出 a=0 b= c=[] d=<nil> e={0}

	// 匿名变量 _
	// 在编码过程中, 可能会遇到没有名称的变量、类型或方法
	fmt.Println("sum =", sum) // 输出 sum = 5
	_, _ = addNumber(&sum)
	fmt.Println("sum =", sum) // 输出 sum = 6
	res, _ := addNumber(&sum)
	fmt.Println("res =", res, "sum =", sum) // 输出 res = 1 sum = 7

}

// addNumber 执行 sum + 1
func addNumber(sum *int) (int, error) {
	// 如果报错就返回 0 错误
	if err := recover(); err != nil {
		return 0, err.(error)
	}

	*sum++

	// 不报错返回 1 nil
	return 1, nil
}
