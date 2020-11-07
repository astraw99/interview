package main

import "fmt"

func main() {
	var d uint = 2
	fmt.Printf("%b\n", ^d) // ^ 一元运算：按位取反
	fmt.Println(len(fmt.Sprintf("%b", ^d)))

	fmt.Println(^uintptr(0)) // ^ 一元运算：按位取反

	fmt.Println(5 ^ 3) // 二元运算：异或xor: 101(5) xor 011(3) = 110(6)

	fmt.Printf("%b\n", uintptr(0))
	fmt.Printf("%b\n", ^uintptr(0))
}
