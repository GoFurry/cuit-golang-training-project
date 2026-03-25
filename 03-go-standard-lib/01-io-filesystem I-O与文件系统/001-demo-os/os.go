package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * title: os demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("cwd =", cwd)
	fmt.Println("tmp dir =", os.TempDir())

	// 环境变量
	_ = os.Setenv("DEMO_KEY", "demo")
	fmt.Println("DEMO_KEY =", os.Getenv("DEMO_KEY"))
	_ = os.Unsetenv("DEMO_KEY")

	// 创建临时文件并写入
	f, err := os.CreateTemp("", "os-demo-*.txt")
	if err != nil {
		fmt.Println("CreateTemp error:", err)
		return
	}
	defer os.Remove(f.Name())
	_, _ = f.WriteString("hello os")
	_ = f.Close()

	// 读取文件内容
	data, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println("ReadFile error:", err)
		return
	}
	fmt.Println("file =", filepath.Base(f.Name()), "content =", string(data))
}
