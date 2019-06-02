package main

import (
	"fmt"
	"runtime"
)

func main() {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
		/*if err := recover(); err != nil {
			fmt.Println(err)
		}*/
	}()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	fmt.Println(runtime.NumGoroutine())

	panic("触发异常")
}
