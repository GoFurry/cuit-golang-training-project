package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

/**
 * title: net/http demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	// 使用 httptest 构造一个本地 HTTP 服务
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/ping" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte("pong"))
	}))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/ping")
	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("status =", resp.StatusCode, "body =", string(body))
}
