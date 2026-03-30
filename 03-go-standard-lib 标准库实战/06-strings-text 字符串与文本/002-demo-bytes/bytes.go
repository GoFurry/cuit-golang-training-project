package main

import (
	"bytes"
	"fmt"
)

/**
 * title: bytes demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	b := []byte("hello bytes")
	fmt.Println("count of 'e':", bytes.Count(b, []byte("e")))

	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteByte(' ')
	buf.WriteString("buffer")
	fmt.Println("buffer:", buf.String())
}
