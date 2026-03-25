package main

import (
	"fmt"
	"regexp"
)

/**
 * title: regexp demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	re := regexp.MustCompile(`(\w+)-(\d+)`)
	match := re.FindStringSubmatch("order-123")
	fmt.Println("match:", match)
}
