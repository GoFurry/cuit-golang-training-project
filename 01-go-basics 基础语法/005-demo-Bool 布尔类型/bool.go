package main

import "fmt"

/**
 * title: Bool demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

func main() {
	// 布尔类型
	//  && 和 || 操作符结合, 有短路行为
	var aVar = 10
	fmt.Println(aVar == 5)  // false
	fmt.Println(aVar == 10) // true
	fmt.Println(aVar != 5)  // true
	fmt.Println(aVar != 10) // false

}
