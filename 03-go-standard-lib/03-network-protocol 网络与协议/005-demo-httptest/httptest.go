package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

/**
 * title: httptest demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	req := httptest.NewRequest(http.MethodGet, "http://example.com/", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	fmt.Println("status =", rec.Code)
	fmt.Println("body =", rec.Body.String())
}
