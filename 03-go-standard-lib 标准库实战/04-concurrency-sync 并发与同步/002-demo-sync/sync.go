package main

import (
	"fmt"
	"sync"
)

/**
 * title: sync demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	var (
		once  sync.Once
		mu    sync.Mutex
		count int
	)

	inc := func() {
		mu.Lock()
		count++
		mu.Unlock()
	}

	for i := 0; i < 3; i++ {
		once.Do(func() { fmt.Println("once") })
		inc()
	}

	fmt.Println("count =", count)
}
