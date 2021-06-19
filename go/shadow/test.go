package main

import "fmt"

func main() {
	x := 100
	func() {
		x := 200 // x will shadow outer x
		//x = 200 // x will override outer x
		fmt.Println(x)
	}()

	fmt.Println(x)
}
