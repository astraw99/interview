package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	// 第一个参数是index
	for v := range s {
		fmt.Println(v)
	}
}
