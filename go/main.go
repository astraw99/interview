package main

import (
	"fmt"
	"runtime"
	"time"
)

// 题目参考：https://zhuanlan.zhihu.com/interview
func main() {
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(2 * i)
		}
	}()

	fmt.Println("------------------")

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(2*i - 1)
		}
	}()

	fmt.Println("num:", runtime.NumGoroutine())
	time.Sleep(time.Second)
}
