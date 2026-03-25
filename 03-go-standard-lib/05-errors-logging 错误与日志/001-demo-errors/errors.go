package main

import (
	"errors"
	"fmt"
)

/**
 * title: errors demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("id=%d not found", e.ID)
}

func main() {
	err := fmt.Errorf("wrap: %w", ErrNotFound)
	fmt.Println("is not found:", errors.Is(err, ErrNotFound))

	err2 := &NotFoundError{ID: 10}
	var nf *NotFoundError
	fmt.Println("as not found:", errors.As(err2, &nf), nf.ID)
}
