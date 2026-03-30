package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

/**
 * title: bufio demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	// Reader
	r := bufio.NewReader(strings.NewReader("line1\nline2\n"))
	line, _ := r.ReadString('\n')
	fmt.Println("ReadString:", strings.TrimSpace(line))

	// Writer
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	_, _ = w.WriteString("hello bufio")
	_ = w.Flush()
	fmt.Println("Writer:", buf.String())
}
