package main

import "fmt"

func doRecover() {
	fmt.Println("recovered =>", recover()) //prints: recovered => <nil>
}

func main() {
	defer func() {
		//doRecover() //panic is not recovered
		fmt.Println(recover()) // 只能直接调用 recover() 才可以捕获 panic
	}()

	panic("not good")
}
