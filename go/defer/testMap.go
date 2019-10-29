package main

import (
	"fmt"
)

func main() {
	mapData := make(map[string]interface{})

	mapData["abc"] = "ABC"
	a := 10

	// defer 匿名函数
	defer func() {
		fmt.Println("======111", EchoInterface(mapData, 1), EchoInterface(a, 1))
	}()

	// 普通 defer
	defer fmt.Println("======222", EchoInterface(mapData, 2), EchoInterface(a, 2))

	mapData["def"] = "DEF"

	a = 20
}

func EchoInterface(i interface{}, index int) interface{} {
	fmt.Println("------", index, i)
	return i
}
