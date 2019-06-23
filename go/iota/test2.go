package main

import (
	"fmt"
)

const (
	aa = iota
	bb = iota
)
const (
	name = "menglu"
	cc   = iota
	dd   = iota
)

func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
}
