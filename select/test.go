package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go testSelectFor(c)

	c <- true
	c <- false
	close(c)

	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Hello, 世界")
}

func testSelectFor(chExit chan bool) {
	for {
		select {
		case v, ok := <-chExit:
			if !ok {
				fmt.Println("close channel 1", v)
				break // can not break for loop
			}

			fmt.Println("ch1 val =", v)
		default:
			fmt.Println("default")
		}
	}

	fmt.Println("exit testSelectFor")
}