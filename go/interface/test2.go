package main

import (
	"fmt"
)

type Person interface {
	Show()
}

type Teacher struct{}

func (tea *Teacher) Show() {

}

func live() Person {
	var tea *Teacher
	return tea
}

func main() {
	v := live()

	// struct 的零值为<nil>，不等于nil
	if v == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Printf("%v, %T", v, v)
		fmt.Println("BBBBBBB")
	}
}
