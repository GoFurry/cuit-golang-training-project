package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * title: Goroutine demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// goroutine 是 Go 的轻量级并发单元

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine-1:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine-2:", i)
			time.Sleep(80 * time.Millisecond)
		}
	}()

	fmt.Println("main waiting...")
	wg.Wait()
	fmt.Println("main done")
}
