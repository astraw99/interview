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

	var sl []string
	var m map[string]string

	// range skip nil slice
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// range skip nil map
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("end")
}
