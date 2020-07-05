package main

import "fmt"

func main() {
	var i = 2

	defer fmt.Println("result =>", func() int { return i * 2 }()) // 运行时求值
	i++
	//prints: result => 2 (not ok if you expected 6)
}
