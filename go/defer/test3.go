package main

import "fmt"

func main() {
	v := c()
	fmt.Println(v)
}

func c() (i int) {

	defer func() { i++ }()

	return 1

}
