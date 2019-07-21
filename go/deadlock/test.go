package main

import "fmt"

func main() {
	c := make(chan int)
	//c <- 99
	go func() {
		c <- 88
	}()
	fmt.Println(<-c)
}
