package main

import "fmt"

/**
 * title: Loop demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// for、break、continue、goto

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45

	// 无限循环, 除非 break 跳出循环
	for {
		sum++
		if sum > 100 {
			fmt.Println(sum) // 101
			break
		}
	}

	// 无限循环, 遇到 goto 跳到对应的标签
	for {
		sum++
		if sum%2 == 0 {
			continue
		}
		fmt.Print(sum, " ") // 103 105 107
		if sum > 105 {
			break
		}
	}

	for {
		sum++
		if sum > 110 {
			goto end
		}
	}
	fmt.Println("before end") // 未被执行
end:
	fmt.Println()
	fmt.Println("after end") // after end
}
