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
	//close(ch) // chan 不关闭则 range 永不退出
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
