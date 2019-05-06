package main

import (
	"fmt"
)

var (
	flag bool
	str  string
)

func foo() {
	flag = true
	str = "setup complete!"
}

func main() {
	go foo()
	for !flag {
		// 按照我们的本意，foo()执行完毕后，flag=true，循环就会退出。
		// 但是其实这个循环会执行N次后才会退出
		fmt.Println("exit now.")
	}
	fmt.Println(str)
}
