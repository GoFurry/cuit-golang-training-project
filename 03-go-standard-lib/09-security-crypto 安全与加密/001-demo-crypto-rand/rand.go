package main

import (
	"crypto/rand"
	"fmt"
)

/**
 * title: crypto/rand demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("rand.Read error:", err)
		return
	}
	fmt.Printf("random bytes: %x\n", b)
}
