package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

/**
 * title: io demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	reader := strings.NewReader("hello io")
	all, _ := io.ReadAll(reader)
	fmt.Println("ReadAll:", string(all))

	// io.Copy
	src := strings.NewReader("copy data")
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, src)
	fmt.Println("Copy:", buf.String())
}
