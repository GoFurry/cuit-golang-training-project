package main

import (
	"database/sql"
	"fmt"
)

/**
 * title: database/sql demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	// database/sql 需要具体驱动，例如 sqlite3 / mysql / postgres
	// 这里演示打开时的错误处理逻辑
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("sql.Open error:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Ping error (need driver):", err)
		return
	}

	fmt.Println("db ready")
}
