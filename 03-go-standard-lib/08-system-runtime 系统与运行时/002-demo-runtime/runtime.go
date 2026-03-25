package main

import (
	"fmt"
	"runtime"
)

/**
 * title: runtime demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	fmt.Println("GOOS =", runtime.GOOS)
	fmt.Println("GOARCH =", runtime.GOARCH)
	fmt.Println("NumCPU =", runtime.NumCPU())
	fmt.Println("NumGoroutine =", runtime.NumGoroutine())

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Alloc =", m.Alloc)
}
