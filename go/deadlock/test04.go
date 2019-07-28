package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // 解决方式1：关闭chan

	// range 一直读取直到chan关闭，否则产生阻塞死锁
	// 解决方式2：开启子协程，主协程sleep等待
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(1e9)
}
