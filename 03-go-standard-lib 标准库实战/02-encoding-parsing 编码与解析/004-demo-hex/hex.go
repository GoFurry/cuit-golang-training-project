package main

import (
	"encoding/hex"
	"fmt"
)

/**
 * title: hex demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	data := []byte("hello hex")
	encoded := hex.EncodeToString(data)
	decoded, _ := hex.DecodeString(encoded)

	fmt.Println("encoded:", encoded)
	fmt.Println("decoded:", string(decoded))
}
