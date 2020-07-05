package main

import "fmt"

func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break // 死循环，一直打印 breaking out...
			break loop
		}
	}
	fmt.Println("out...")

	fmt.Println("out")

}
