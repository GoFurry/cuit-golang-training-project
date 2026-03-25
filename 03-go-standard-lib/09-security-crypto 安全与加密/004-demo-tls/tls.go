package main

import (
	"crypto/tls"
	"fmt"
)

/**
 * title: tls demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	cfg := &tls.Config{MinVersion: tls.VersionTLS12}
	fmt.Println("min version =", cfg.MinVersion)
	fmt.Println("cipher suites =", len(tls.CipherSuites()))
}
