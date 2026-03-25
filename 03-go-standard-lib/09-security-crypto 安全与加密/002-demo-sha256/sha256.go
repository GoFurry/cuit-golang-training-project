package main

import (
	"crypto/sha256"
	"fmt"
)

/**
 * title: sha256 demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	sum := sha256.Sum256([]byte("hello sha256"))
	fmt.Printf("sha256: %x\n", sum)
}
