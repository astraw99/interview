package main

import (
	"fmt"
)

func main() {
	var a interface{} = "10"
	b, ok := a.(int)
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}

	fmt.Println(a)
	fmt.Println(b)
}
