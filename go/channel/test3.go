package main

import (
	"fmt"
	"time"
)

var endRunning = make(chan int, 1)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		endRunning <- 1
		fmt.Println("send to endRunning")
	}()

	close(endRunning)

	select {
	case v, ok := <-endRunning:
		if !ok {
			fmt.Println("chan closed")
		} else {
			fmt.Printf("receive endRunning: %v, ok-%v\n", v, ok)
		}
		/*default:
		fmt.Println("default wait")*/
	}

	fmt.Println("finish")

	return
}
