package main

import (
	"errors"
	"fmt"
)

/**
 * title: ErrorHandling demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// Go 使用 error 接口进行错误处理
// 通过 errors.Is/As 判断错误类型或解包错误链

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found, id=%d", e.Resource, e.ID)
}

func query(id int) error {
	if id == 0 {
		return fmt.Errorf("query id=%d: %w", id, ErrNotFound)
	}
	return nil
}

func loadUser(id int) error {
	if id <= 0 {
		return &NotFoundError{Resource: "user", ID: id}
	}
	return nil
}

func main() {
	if err := query(0); err != nil {
		fmt.Println("query error:", err)
		if errors.Is(err, ErrNotFound) {
			fmt.Println("matched ErrNotFound")
		}
	}

	if err := loadUser(-1); err != nil {
		fmt.Println("loadUser error:", err)
		var nf *NotFoundError
		if errors.As(err, &nf) {
			fmt.Println("resource:", nf.Resource, "id:", nf.ID)
		}
	}
}
