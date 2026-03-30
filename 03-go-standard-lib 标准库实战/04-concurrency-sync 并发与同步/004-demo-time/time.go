package main

import (
	"fmt"
	"time"
)

/**
 * title: time demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	now := time.Now()
	fmt.Println("now =", now.Format(time.RFC3339))

	parsed, _ := time.Parse("2006-01-02", "2026-03-20")
	fmt.Println("parsed =", parsed.Format(time.RFC3339))

	<-time.After(50 * time.Millisecond)
	fmt.Println("after 50ms")
}
