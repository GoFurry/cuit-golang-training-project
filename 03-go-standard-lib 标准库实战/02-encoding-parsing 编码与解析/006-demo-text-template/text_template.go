package main

import (
	"bytes"
	"fmt"
	"text/template"
)

/**
 * title: text/template demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

type Profile struct {
	Name string
	Age  int
}

func main() {
	tpl := "Hello {{.Name}}, age {{.Age}}"
	t, _ := template.New("t").Parse(tpl)

	var buf bytes.Buffer
	_ = t.Execute(&buf, Profile{Name: "Alice", Age: 20})
	fmt.Println(buf.String())
}
