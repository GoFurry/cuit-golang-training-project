package main

import (
	"encoding/json"
	"fmt"
)

/**
 * title: json demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{Name: "Alice", Age: 20}

	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("marshal:\n", string(data))

	var out User
	_ = json.Unmarshal(data, &out)
	fmt.Println("unmarshal:", out)
}
