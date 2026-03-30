package main

import (
	"fmt"
	"path/filepath"
)

/**
 * title: filepath demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	p := filepath.Join("/tmp", "demo", "file.txt")
	fmt.Println("join =", p)
	fmt.Println("clean =", filepath.Clean("/tmp/../tmp/demo//file.txt"))
	fmt.Println("base =", filepath.Base(p))
	fmt.Println("dir =", filepath.Dir(p))
	fmt.Println("ext =", filepath.Ext(p))
}
