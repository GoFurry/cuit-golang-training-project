package main

import (
	"encoding/xml"
	"fmt"
)

/**
 * title: xml demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
}

func main() {
	b := Book{Title: "Go", Author: "Tom"}
	data, _ := xml.MarshalIndent(b, "", "  ")
	fmt.Println("xml:\n", string(data))

	var out Book
	_ = xml.Unmarshal(data, &out)
	fmt.Println("unmarshal:", out.Title, out.Author)
}
