package main

import (
	"fmt"
	"time"
)

var endRunning = make(chan bool, 1)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		endRunning <- true
		fmt.Println("send to endRunning")
	}()

	<-endRunning
	fmt.Println("receive endRunning")

	return
}
