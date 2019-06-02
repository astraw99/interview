package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)

	// array 传值会copy一份
	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	y := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(&y)

	// array 传指针会同步更新
	fmt.Println(y)
}
