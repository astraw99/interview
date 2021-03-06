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

	strFloat := ".2abc"
	float, err := strconv.ParseFloat(strFloat, 64)

	// 0 strconv.Atoi: parsing "abc2": invalid syntax
	fmt.Println(float, err)

	// 0 strconv.ParseInt: parsing "": invalid syntax
	num, err := strconv.ParseInt("", 10, 64)
	fmt.Println(num, err)
}
