package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // 超过最大容量，阻塞main协程，产生deadlock

	for v := range ch {
		fmt.Println(v)
	}
}
