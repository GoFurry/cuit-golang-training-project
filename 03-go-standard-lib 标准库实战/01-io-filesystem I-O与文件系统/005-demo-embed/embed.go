package main

import (
	"embed"
	"fmt"
)

/**
 * title: embed demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

//go:embed data.txt
var data string

func main() {
	fmt.Println("embed data:", data)
}
