package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	//<-ch // 读取空chan产生死锁

	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("chan no data")
	}
}
