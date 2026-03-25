package main

import (
	"fmt"
	"sync/atomic"
)

/**
 * title: sync/atomic demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	var n int64
	atomic.AddInt64(&n, 1)
	atomic.AddInt64(&n, 2)
	fmt.Println("n =", atomic.LoadInt64(&n))
}
