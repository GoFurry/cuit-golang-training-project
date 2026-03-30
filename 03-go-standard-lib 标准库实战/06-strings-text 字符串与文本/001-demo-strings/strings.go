package main

import (
	"fmt"
	"strings"
)

/**
 * title: strings demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	s := "go,lang,strings"
	fmt.Println("contains:", strings.Contains(s, "lang"))
	fmt.Println("split:", strings.Split(s, ","))
	fmt.Println("join:", strings.Join([]string{"go", "lang"}, "-"))
	fmt.Println("replace:", strings.ReplaceAll(s, ",", "|"))
}
