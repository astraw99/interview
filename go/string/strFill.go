package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "12345"
	ss := fmt.Sprintf("%032s", "") // 表示以 0 填充 32 位字符串
	a = fmt.Sprintf("%.32s", a+ss) // 宽度.精度 字符串精度表示最大字符数

	fmt.Println(ss, a)

	b := strings.Repeat("=", 20) // 字符串重复填充
	fmt.Println(b)
}
