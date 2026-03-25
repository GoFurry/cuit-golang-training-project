package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
)

/**
 * title: csv demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"name", "score"})
	_ = w.Write([]string{"Tom", "90"})
	_ = w.Write([]string{"Lucy", "88"})
	w.Flush()

	fmt.Println("csv data:\n", buf.String())

	r := csv.NewReader(strings.NewReader(buf.String()))
	records, _ := r.ReadAll()
	fmt.Println("records:", records)
}
