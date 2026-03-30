package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * title: utf8 demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	s := "你好go"
	fmt.Println("byte len:", len(s))
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	r, size := utf8.DecodeRuneInString(s)
	fmt.Printf("first rune: %c size=%d\n", r, size)
}
