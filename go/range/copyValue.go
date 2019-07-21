package main

import (
	"fmt"
	"strconv"
)

type Foo struct {
	Bar string
}

func main() {
	list := []Foo{{"bar1"}, {"bar2"}, {"bar3"}}
	for i, v := range list {
		v.Bar = "change" + strconv.Itoa(i)
		//fmt.Println(v)
	}

	a := []int{1, 2, 3}
	for _, v := range a {
		// for range 中的 v 只会声明初始化一次，不会每次循环都声明初始化
		fmt.Println(&v, v)
	}

	fmt.Println(list)
}
