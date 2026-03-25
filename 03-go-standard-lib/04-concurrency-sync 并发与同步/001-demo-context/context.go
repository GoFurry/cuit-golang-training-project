package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * title: context demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("ctx done:", ctx.Err())
	}
}
