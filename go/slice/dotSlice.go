package main

import "fmt"

func main() {
	printDot(1, 2)
}

func printDot(a ...int) {
	fmt.Printf("%T\n", a) // []int

	for _, v := range a {
		fmt.Println(v)
	}
}
