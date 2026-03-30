package main

import (
	"log/slog"
	"os"
)

/**
 * title: slog demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("login", "user", "alice", "ip", "127.0.0.1")
}
