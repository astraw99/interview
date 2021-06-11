package main

import (
	"fmt"
)

func main() {
	if !test1() || !test2() && test3() {
		fmt.Println("111111111")
		return
	}

	fmt.Println("22222222")
}

func test1() bool {
	fmt.Println("test1")
	return true
}

func test2() bool {
	fmt.Println("test2")
	return true
}

func test3() bool {
	fmt.Println("test3")
	return true
}
