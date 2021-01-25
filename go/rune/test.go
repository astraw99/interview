package main

import "fmt"

func main()  {
	str := "Go 爱好者"

	for i, c := range str {
		// 索引是 Unicode 编码值第一个字节对应索引值
		// c 为 rune 类型(int32)
		fmt.Printf("%d, %q, %T, %v\n", i, c, c, []byte(string(c)))
	}
}
