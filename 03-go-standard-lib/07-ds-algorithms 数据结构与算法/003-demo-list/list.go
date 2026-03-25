package main

import (
	"container/list"
	"fmt"
)

/**
 * title: list demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	l := list.New()
	l.PushBack("A")
	l.PushBack("B")
	l.PushFront("Start")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("val =", e.Value)
	}
}
