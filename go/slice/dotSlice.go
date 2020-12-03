package main

import "fmt"

func main() {
	printDot(1, 2)
}

func printDot(a ...int) {
	fmt.Printf("%T is slice\n", a) // []int

	for _, v := range a {
		fmt.Println(v)
	}

	var nums = [...]int{1, 2, 3} // 省略号可自动计算数组长度
	var numsArr = [3]int{1, 2, 3}
	fmt.Printf("%T is array\n", nums) // [3]int
	fmt.Printf("%T is array\n", numsArr) // [3]int
}
