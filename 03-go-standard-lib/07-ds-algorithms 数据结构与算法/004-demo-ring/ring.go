package main

import (
	"container/ring"
	"fmt"
)

/**
 * title: ring demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	r := ring.New(3)
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}

	r.Do(func(v interface{}) {
		fmt.Println("val =", v)
	})
}
