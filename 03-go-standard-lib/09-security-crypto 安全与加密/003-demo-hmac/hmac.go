package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

/**
 * title: hmac demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	key := []byte("secret")
	m := hmac.New(sha256.New, key)
	m.Write([]byte("message"))
	sum := m.Sum(nil)
	fmt.Printf("hmac: %x\n", sum)
}
