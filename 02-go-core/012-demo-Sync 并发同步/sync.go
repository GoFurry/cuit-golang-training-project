package main

import (
	"fmt"
	"sync"
)

/**
 * title: Sync demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// sync 包提供常用的并发同步工具，如 WaitGroup、Mutex

func main() {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
	)

	// 启动 1000 个 goroutine 安全地累加
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("counter =", counter)
}
