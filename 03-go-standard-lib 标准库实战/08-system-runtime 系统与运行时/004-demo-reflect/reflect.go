package main

import (
	"fmt"
	"reflect"
)

/**
 * title: reflect demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	var x int = 10
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("type =", t.Name(), "value =", v.Int())

	// 通过指针修改值
	pv := reflect.ValueOf(&x).Elem()
	pv.SetInt(99)
	fmt.Println("x =", x)
}
