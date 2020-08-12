package main

import (
	"fmt"
	"runtime"
	// "time"
)

var count = 15

func main() {
	runtime.GOMAXPROCS(1)

	// 下面的1和2交换位置，结果却不同，原因是？
	// 1.偶数
	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2 * i)
			runtime.Gosched()
		}
	}()

	// 2.奇数
	go func() {
		for i := 1; i < count; i++ {
			fmt.Println(2*i - 1)
			runtime.Gosched()
		}
	}()

	var input string
	fmt.Scanln(&input)
}
