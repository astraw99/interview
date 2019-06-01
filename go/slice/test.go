package main

import "fmt"

func main() {
	slice := make([]string, 10)

	// append追加后，len + 1
	slice = append(slice, "a")

	fmt.Println(len(slice), slice)
}
