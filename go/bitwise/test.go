package main

import "fmt"

func main() {
	var d uint = 2
	fmt.Printf("%b\n", ^d)
	fmt.Println(len(fmt.Sprintf("%b", ^d)))
}
