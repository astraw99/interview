package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	// 只有一个返回值：则第一个参数是index
	for i := range s {
		fmt.Println(i)
	}

	// 两个返回值
	for i, v := range s {
		fmt.Println(i, v)
	}
}
