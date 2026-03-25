package main

import (
	"fmt"
	"sort"
)

/**
 * title: sort demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

func main() {
	nums := []int{5, 2, 9, 1}
	sort.Ints(nums)
	fmt.Println("sorted:", nums)

	people := []struct {
		Name string
		Age  int
	}{{"Tom", 18}, {"Lucy", 20}, {"Bob", 17}}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("people:", people)
}
