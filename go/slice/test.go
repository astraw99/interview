package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	slice := make([]string, 10)

	// append追加后，len + 1
	slice = append(slice, "a")

	fmt.Println(len(slice), slice)
}
