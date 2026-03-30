package main

import (
	"flag"
	"fmt"
)

/**
 * title: flag demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	name := flag.String("name", "gopher", "user name")
	age := flag.Int("age", 18, "user age")
	flag.Parse()

	fmt.Println("name =", *name, "age =", *age)
}
