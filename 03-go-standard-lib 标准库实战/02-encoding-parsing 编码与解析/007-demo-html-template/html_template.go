package main

import (
	"bytes"
	"fmt"
	"html/template"
)

/**
 * title: html/template demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	tpl := "<div>{{.}}</div>"
	t, _ := template.New("t").Parse(tpl)

	var buf bytes.Buffer
	_ = t.Execute(&buf, "<b>safe?</b>")
	fmt.Println(buf.String())
}
