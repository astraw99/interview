package main

import (
	"fmt"
)

func main() {
	var x string = nil // string 零值不是 nil，而是 ""
	//var x string = "" // string 零值不是 nil，而是 ""
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
