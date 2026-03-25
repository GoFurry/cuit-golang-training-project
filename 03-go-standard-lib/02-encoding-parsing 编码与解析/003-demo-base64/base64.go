package main

import (
	"encoding/base64"
	"fmt"
)

/**
 * title: base64 demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	plain := "hello base64"
	encoded := base64.StdEncoding.EncodeToString([]byte(plain))
	decoded, _ := base64.StdEncoding.DecodeString(encoded)

	fmt.Println("encoded:", encoded)
	fmt.Println("decoded:", string(decoded))
}
