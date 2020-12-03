package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 1
	ch <- 3

	//close(ch) // range不等到信道关闭，不会结束读取，阻塞当前的goroutine，导致死锁
	for c := range ch {
		fmt.Println(c)
		/*if len(ch) == 0 {
			break
		}*/
	}
}
