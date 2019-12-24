package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//str := "abcgegegjijieggekgoege454g5e45g4e5"

	//slice := make([]string, 0) // 空 slice 不会走进 for range
	slice := make([]string, 5)
	rand.Seed(time.Now().Unix())

	for i, v := range slice {
		fmt.Println("index:", i)
		fmt.Println("value:", v)
		slice[i] = fmt.Sprintf("%d", rand.Intn(5))
	}

	fmt.Println(slice)
}
