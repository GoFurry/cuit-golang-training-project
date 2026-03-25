package main

import "fmt"

/**
 * title: Map demo
 * author: 福狼
 * create_time: 2026.03.20
 * update_time: 2026.03.20
 */

// map 是无序的键值对集合
// 未初始化的 map 是 nil，写入会 panic

func main() {
	// nil map
	var nilMap map[string]int
	fmt.Println("nilMap == nil:", nilMap == nil)

	// make 初始化
	ages := make(map[string]int)
	ages["Tom"] = 18
	ages["Lucy"] = 20
	fmt.Println("ages =", ages)

	// 字面量初始化
	scores := map[string]int{
		"math":  90,
		"music": 88,
	}
	fmt.Println("scores =", scores)

	// 读取与 ok 判断
	val, ok := scores["math"]
	fmt.Println("math:", val, "ok:", ok)

	val, ok = scores["english"]
	fmt.Println("english:", val, "ok:", ok)

	// 删除键
	delete(scores, "music")
	fmt.Println("scores after delete =", scores)

	// 遍历（map 遍历顺序是随机的）
	for k, v := range ages {
		fmt.Printf("%s => %d\n", k, v)
	}
}
