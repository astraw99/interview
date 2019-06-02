package main

import "fmt"

func main() {
	x := "text"
	// 字符串索引值是字节ASCII值
	fmt.Println(x[0])

	// 字符串是常量不能变，可以转换为byte[]再操作
	//x[0] = 'T'

	xbytes := []byte(x)
	xbytes[0] = 'T'

	x = string(xbytes)

	fmt.Println(x)
}
