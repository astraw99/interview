package main

import "fmt"

func main() {
	x := 100
	func() {
		x := 200 // 会 shadow 外层的x
		fmt.Println(x)
	}()

	fmt.Println(x)
}
