package main

import (
	"fmt"
	"net/url"
)

/**
 * title: net/url demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	u, _ := url.Parse("https://example.com/search?q=golang&page=1")
	fmt.Println("host =", u.Host)
	fmt.Println("path =", u.Path)
	fmt.Println("q =", u.Query().Get("q"))

	ref, _ := url.Parse("/docs")
	fmt.Println("resolve =", u.ResolveReference(ref))
}
