package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
 * title: net demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	c1, c2 := net.Pipe()
	defer c1.Close()
	defer c2.Close()

	go func() {
		_, _ = c1.Write([]byte("hello net"))
	}()

	reader := bufio.NewReader(c2)
	msg, _ := reader.ReadString('t')
	fmt.Println("recv:", msg)
}
