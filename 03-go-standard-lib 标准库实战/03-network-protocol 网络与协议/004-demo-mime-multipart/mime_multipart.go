package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
)

/**
 * title: mime/multipart demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	_ = writer.WriteField("name", "alice")
	file, _ := writer.CreateFormFile("file", "demo.txt")
	_, _ = file.Write([]byte("file content"))
	_ = writer.Close()

	fmt.Println("content-type =", writer.FormDataContentType())
	fmt.Println("body size =", buf.Len())
}
