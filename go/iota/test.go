package main

import "fmt"

const (
	a        = iota
	b        = 100
	c uint64 = 1 << 63
	d
	e = iota
)

func main() {
	fmt.Println(a, b, c, d, e)
}
