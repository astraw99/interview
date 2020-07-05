package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("Processed:", m)
			//time.Sleep(1 * time.Second) // 模拟需要长时间运行的操作
			fmt.Println("process ok")
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" // 不会被接收处理，为什么？

	fmt.Println("ok")
}
