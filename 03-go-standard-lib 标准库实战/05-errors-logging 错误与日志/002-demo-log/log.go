package main

import "log"

/**
 * title: log demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("hello log")
}
