package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int // 未初始化，值为 nil
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	fmt.Println("Result: ", <-ch)
	time.Sleep(2 * time.Second)
}
