package main

import (
	"fmt"
	"strings"
)

/**
 * title: String demo
 * author: 福狼
 * create_time: 2025.12.03
 * update_time: 2025.12.03
 */

// Golang 使用 UTF-8
// UTF-8 字符占用 1-4 个字节, ASCII 表上的占 1 个, 其余占 2-4 个
// 字符串是一种值类型, 且值不可变, 即字符串是字节的定长数组

func main() {
	// 使用""定义字符串 通过\使用转义字符
	// \n 换行符
	// \r 回车符
	// \t tab 键
	// \u 或 \U Unicode 字符
	// \\ 反斜杠自身
	var str = "成信大Go语言教程\nCUIT Golang training guide"
	fmt.Println(str)
	// ``不会转义任何东西 纯字符串
	var strA = `成信大Go语言教程\nCUIT Golang training guide`
	fmt.Println(strA)
	// ''定义单字符
	var strB = '成'
	fmt.Println(strB)         // 字节码 25104
	fmt.Println(string(strB)) // 字符 成

	// 运算符会比较字符串自然编码的顺序
	fmt.Println("abcde" > "abcd")  // true
	fmt.Println("abcde" > "abcdd") // true
	fmt.Println("abcde" > "abcfe") // false

	// 字节可以用索引获取(不是字符)
	fmt.Println(str[0])           // 230
	fmt.Println(str[len(str)-1])  //100
	fmt.Println(str[0:3])         // 成
	fmt.Println(str[len(str)-1:]) // e

	// 字符串拼接
	// 性能比较 Builder > Join > + > Sprintf
	// Sprintf 底层涉及反射、格式化解析、临时缓冲区, 性能损耗最大
	// + 每次都会创建新字符串, 内存分配频繁
	// Join 先计算所有字符串总长度, 一次性分配内存, 再拷贝数据
	// Builder 底层是[]byte缓冲区, 拼接时仅向缓冲区写入, 无多余内存拷贝
	var s1, s2 string = "Hello", "Golang"
	fmt.Println(s1 + s2)                     // + 拼接 HelloGolang
	fmt.Println(fmt.Sprintf("%s%s", s1, s2)) // fmt包S开头方法拼接 HelloGolang
	strings.Join([]string{s1, s2}, "")       // strings.Join 拼接 HelloGolang
	// strings.Builder 拼接 HelloGolang
	{
		var builder strings.Builder
		for _, arg := range []string{s1, s2} {
			builder.WriteString(arg)
		}
		fmt.Println(builder.String())
	}

	// 书写 Unicode 字符, 4 字节用 \u, 8 字节用 \U
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
