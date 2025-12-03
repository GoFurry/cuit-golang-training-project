package main

//    包声明: main 包是 Go 程序的入口包, 包含 main 函数的源文件必须声明为 main 包
//    只有 main 包的程序才能编译为可执行文件, 其他包通常作为库被导入使用

/**
 * title: HelloWorld demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

import "fmt"

//    导入标准库: fmt (format) 提供格式化输入/输出功能
//    包含 Print/Println/Printf/Scan/Scanln 等常用函数

// main函数: Go程序的唯一入口函数, 程序启动时自动执行, 无参数、无返回值
// 一个Go程序只能有一个main函数
func main() {
	// Print 将内容输出到标准输出(控制台), 不会自动换行
	fmt.Print("Hello World ") // 输出"Hello World ", 末尾空格保留，无换行
	fmt.Print("Hello World ") // 继续在当前行输出, 最终该行显示"Hello World Hello World "

	// Println 输出内容后自动追加一个换行符\n
	fmt.Println("Hello Golang") // 输出"Hello Golang"并换行
	fmt.Println("Hello Golang") // 在下一行输出"Hello Golang"并换行

	// Printf 按指定格式字符串输出内容
	// 第一个参数是格式字符串, 后续参数为占位符对应的值
	// 常用占位符：
	//   %s - 字符串类型
	//   %v - 任意类型的默认格式
	//   %d - 整数类型
	//   %f - 浮点类型
	//   \n - 换行符
	fmt.Printf("%s %v ovo\n", "Hello", "Gopher")
	// 最终输出: Hello Gopher ovo
}
