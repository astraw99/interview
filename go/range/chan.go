package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	//close(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for v := range ch {
			fmt.Println(v)
			if len(ch) == 0 {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
