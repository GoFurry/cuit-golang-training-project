package main

import "fmt"

/**
 * title: Channel demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// channel 用于 goroutine 之间通信
// 未缓冲通道需要发送与接收同时发生

func main() {
	// 无缓冲通道
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("recv:", v)
	}

	// 有缓冲通道
	buf := make(chan string, 2)
	buf <- "A"
	buf <- "B"
	fmt.Println("buf1:", <-buf)
	fmt.Println("buf2:", <-buf)
}
