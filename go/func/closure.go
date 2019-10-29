package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			fmt.Println("------")
			return name + suffix
		}

		fmt.Println("=======")
		return name
	}
}

func main() {
	//判断字符串 以bmp结尾
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("test"))
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}
