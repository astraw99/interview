package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday().String()
	fmt.Printf("%T\n%v\n", day, day)

	fmt.Println(time.Now().Format("15:04"))

	fmt.Println(time.Now().Format("15:04") >= "03:07")

	t, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 11:20"+":00")

	fmt.Println(t, err)
	fmt.Println(t.Add(time.Duration(200 * 1e9)))

	fmt.Println(time.Now().Add(10 * time.Hour))

	fmt.Println(time.Now().Unix())
}
