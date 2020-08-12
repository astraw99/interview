package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	fmt.Println("ch:", ch) // 未初始化，值为 nil

	for {
		select {
		case v, ok := <-ch: // nil chan 不会触发 select
			fmt.Println(v, ok)
		default:
			fmt.Println("default")
		}

		time.Sleep(1 * time.Second)
	}

	//fmt.Println("Result: ", <-ch)
	time.Sleep(2 * time.Second)
}
