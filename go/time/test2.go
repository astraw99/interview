package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			v := <-ch
			fmt.Println("go func v:", v)
		}
	}(ch)

	tick := time.NewTicker(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case ch <- i:
		}

		// ticker 单独控制
		select {
		case <-tick.C:
			fmt.Println("tick timeout:", i)
		//default:
		//	fmt.Println("tick default")
		}

		time.Sleep(200 * time.Millisecond)
	}

	close(ch)
	tick.Stop()


	fmt.Println(time.December.String())
}
