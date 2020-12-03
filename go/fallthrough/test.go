package main

import "fmt"

func main() {
	// switch 默认匹配 true
	switch {
	case true:
		fmt.Println("The integer was <= 1")
		fallthrough
	case false:
		fmt.Println("The integer was <= 2")
		//break (break 可省略)
	default:
		fmt.Println("default case")
	}
}
