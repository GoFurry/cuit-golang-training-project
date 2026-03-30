package main

import "fmt"

/**
 * title: Interface demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// 接口是一组方法集合，Go 通过“隐式实现”来满足接口

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": wang"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": miao"
}

func say(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	var s Speaker
	s = Dog{Name: "旺财"}
	say(s)

	s = Cat{Name: "咪咪"}
	say(s)

	// 空接口可以接收任意类型
	var anything interface{}
	anything = 42
	fmt.Println("anything =", anything)

	// 类型断言
	v, ok := anything.(int)
	fmt.Println("assert int:", v, ok)

	// 类型选择
	switch t := anything.(type) {
	case int:
		fmt.Println("type is int:", t)
	case string:
		fmt.Println("type is string:", t)
	default:
		fmt.Println("unknown type")
	}
}
