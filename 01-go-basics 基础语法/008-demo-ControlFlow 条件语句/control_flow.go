package main

import "fmt"

/**
 * title: Control flow demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// if 语句和 switch语句

func main() {
	// if 语句
	conditionTrue, conditionFalse := true, false
	if conditionTrue {
		fmt.Println("if true")
	}
	// 输出 if true

	// else
	if conditionFalse {
		fmt.Println("if true")
	} else {
		fmt.Println("if false")
	}
	// 输出 if false

	// else if
	if conditionFalse {
		fmt.Println("if true")
	} else if conditionFalse {
		fmt.Println("if false")
	} else {
		fmt.Println("else if true")
	}
	// 输出 else if true

	// if 内可以用表达式
	if conditionTrue != true {
		fmt.Println("conditionTrue if false")
	} else {
		fmt.Println("conditionTrue if true")
	}
	// 输出 conditionTrue if true

	// if 常用作处理错误
	if _, err := GetNumber(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("err is nil")
	}
	// 输出 err is nil

	// switch
	switch conditionTrue {
	case true:
		fmt.Println(1)
	case false:
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
	// 输出 1

	// 一分支多值
	switch conditionFalse {
	case true, false:
		fmt.Println(1)
	default:
		fmt.Println(0)
	}
	// 输出 1

	// 分钟表达式
	switch {
	case conditionTrue == true:
		fmt.Println(1)
	default:
		fmt.Println(3)
	}
	// 输出 1

	// fallthrough
	switch {
	case conditionTrue == true:
		fmt.Printf("hello ")
		fallthrough
	case conditionTrue == false:
		fmt.Println("world")
	}
	// 输出 hello world
}

func GetNumber() (int, error) {
	return 1, nil
}
