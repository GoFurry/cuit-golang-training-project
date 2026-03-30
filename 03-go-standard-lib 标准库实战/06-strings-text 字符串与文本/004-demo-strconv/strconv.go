package main

import (
	"fmt"
	"strconv"
)

/**
 * title: strconv demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	i, _ := strconv.Atoi("123")
	fmt.Println("atoi:", i)
	fmt.Println("itoa:", strconv.Itoa(456))

	b, _ := strconv.ParseBool("true")
	fmt.Println("parse bool:", b)

	f := strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Println("format float:", f)
}
