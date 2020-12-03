package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	intChan := make(chan int, 1)

	// ⼀秒后关闭通道
	time.AfterFunc(time.Second, func() {
		// new goroutine
		fmt.Println(runtime.NumCgoCall(), runtime.NumCPU(), runtime.NumGoroutine())
		close(intChan)
	})

	fmt.Println("==============")


	// 带缓存的channel实际上是一个阻塞队列
	// 对队列的写总发生在队尾，队列满时写协程阻塞
	// 对队列的读总发生在队首，队列空时读协程阻塞
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}

		fmt.Println("The candidate case is selected.")
	}

	fmt.Println(runtime.NumCgoCall(), runtime.NumCPU(), runtime.NumGoroutine())

	fmt.Println(cap(intChan), len(intChan))
}
