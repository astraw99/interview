package main

import "fmt"

func main() {
	// 0000 0001 => -000 0010 = -2
	fmt.Println(^1)

	// 0000 0010 => -000 0011 = -3
	fmt.Println(^2)

	// 0000 0101 => -000 0110 = -6
	fmt.Println(^80)

	fmt.Printf("%08b\n", 8)
	fmt.Printf("%08b\n", ^8)

	// 0000 0001 + 0000 0010 = 0000 0011 = 3
	fmt.Println(1 ^ 2)
}
