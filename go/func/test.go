package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Println(f(1), "-")
	fmt.Println(f(20), "-")
	fmt.Println(f(300), "-")
}

func Adder() func(int) int {
	//var x int
	x := 0
	return func(delta int) int {
		fmt.Println("*x -", &x) // 内存地址不变
		x += delta
		return x
	}
}
