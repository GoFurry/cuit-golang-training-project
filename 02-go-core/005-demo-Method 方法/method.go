package main

import "fmt"

/**
 * title: Method demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 方法是带接收者的函数
// 值接收者会拷贝，指针接收者可修改原对象

type Counter struct {
	Value int
}

// Add 值接收者，不会修改原对象
func (c Counter) Add(n int) {
	c.Value += n
	fmt.Println("Add inside =", c.Value)
}

// AddInPlace 指针接收者，会修改原对象
func (c *Counter) AddInPlace(n int) {
	c.Value += n
}

// Reset 清零
func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	c := Counter{Value: 10}
	c.Add(5) // 值接收者：不会影响外部
	fmt.Println("after Add =", c.Value)

	c.AddInPlace(5)
	fmt.Println("after AddInPlace =", c.Value)

	c.Reset()
	fmt.Println("after Reset =", c.Value)
}
