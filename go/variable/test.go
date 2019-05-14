package main

import (
	"fmt"
	"os"
)

var a, b, c int = 1, 2, 3

const i uint = 100

//var a int, b int, c int // 不能这样写

//str := "china" // 变量初始化简写只能在函数内

/*type MyFile struct {
	F *os.File
}*/

func main() {
	str := "china"

	fmt.Println(str)

	const a int = 1

	//var mf MyFile // := 不能赋值给结构体字段
	//var err error // 没有新变量不能使用 :=

	f, err := os.Open(".gitignore")
	fmt.Println(f, err)
}
