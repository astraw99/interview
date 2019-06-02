package main

import "fmt"

func main() {
	m := make(map[int]int, 100)

	m[0] = 5
	m[1] = 10

	// map由系统进行内存分配，不支持cap
	//fmt.Println(cap(m))
	fmt.Println(len(m))

	s := new([]string)
	// len & cap 仅支持 pointer of array
	//fmt.Println(len(s))
	fmt.Println(len(*s))
	fmt.Println(cap(*s))
}
