package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1

	fmt.Println(<-ch)

	/*ch <- 1 // 无缓冲在此写入数据，却没有读数据，阻塞住

	fmt.Println(<-ch) // 被上面阻塞，无法被执行到*/

	/*go func() {
		ch <- 1 // 开启子goroutine写入数据
	}()

	fmt.Println(<-ch) // 阻塞住，一旦ch有数据，则读取成功*/
}
