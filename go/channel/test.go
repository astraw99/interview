package main

import (
	"fmt"
	"runtime"
	"time"
)

var count = 15

func ping(c chan<- int) {
	for i := 1; i < count; i++ {
		c <- 2*i - 1
	}
}
func pong(c chan<- int) {
	for i := 1; i < count; i++ {
		c <- 2 * i
	}
}

func echo(ch <-chan int) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	ch := make(chan int)
	go ping(ch)
	go pong(ch)
	go echo(ch)

	/*for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 50)
	}*/

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	// 监听到键盘输入有换行符后程序退出
	var input string
	fmt.Scanln(&input)
}
