package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	slice := make([]string, 10)
	// append追加后，len + 1
	slice = append(slice, "a")
	fmt.Println(len(slice), slice)


	// initSlice的扩容轨迹是1-2-4。而slice只是引用类型，所以x和y只是copy了initSlice的指针，他们三个都是指向同一个底层数组，所以最后第四个坑被y给覆盖了
	initSlice := []int{1}
	// 进行扩容到2
	initSlice = append(initSlice, 2)
	// 进行扩容到4
	initSlice = append(initSlice, 3)
	x := append(initSlice, 4)
	y := append(initSlice, 5)
	fmt.Println(initSlice, x, y)
	fmt.Printf("%p, %p, %p", initSlice, x, y)
}
