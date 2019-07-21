package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "sent result")
			case <-done:
				//case v, ok := <-done:
				fmt.Println(idx, "exiting")
				//fmt.Println(idx, "exiting", v, ok)
			}
		}(i)
	}

	//get first result
	close(done)
	fmt.Println("result:", <-ch)
	//do other work
	time.Sleep(3 * time.Second)
}
