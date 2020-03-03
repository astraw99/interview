package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "abc2"
	pageNumInt, err := strconv.Atoi(str)

	// 0 strconv.Atoi: parsing "abc2": invalid syntax
	fmt.Println(pageNumInt, err)
}
