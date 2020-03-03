package main

import "fmt"

func main() {
	demo := map[string]string{
		"a": "aaa",
	}
	fmt.Println(demo["a"])

	_, ok := demo["a"] //判断a是否存在
	fmt.Println(ok)    //true

	v, ok2 := demo["b"]
	fmt.Printf("%T\n", v)
	fmt.Println(len(v), v, ok2) //false
}
