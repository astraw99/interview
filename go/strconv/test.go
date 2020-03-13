package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "115.2"
	// [核心] Atoi 转浮点数会报错，返回0
	pageNumInt, err := strconv.Atoi(str)

	// 0 strconv.Atoi: parsing "115.2": invalid syntax
	fmt.Println(pageNumInt, err)

	strFloat := "abc2"
	float, err := strconv.ParseFloat(strFloat, 64)

	// 0 strconv.Atoi: parsing "abc2": invalid syntax
	fmt.Println(float, err)
}
