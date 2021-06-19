package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("func err1:", test1())
	//fmt.Println("func err2:", test2())
	fmt.Println("func err3:", test3())
	fmt.Println("func err4:", test4())
	fmt.Println("func err5:", test5())
}

func test1() error {
	var err error

	defer func() {
		fmt.Println("defer err1:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return err
	}

	return nil
}

func test2() (err error) {

	defer func() {
		fmt.Println("defer err2:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return // return without err will compilation error
	}

	return
}

func test3() (err error) {

	defer func() {
		fmt.Println("defer err3:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return err
	}

	return
}

func test4() (err error) {

	defer func() {
		fmt.Println("defer err4:", err)
	}()

	if _, err := os.Open("xxx"); err == nil {
		if err := json.Unmarshal([]byte("{}"), &struct{}{}); err == nil {
			fmt.Println("OK")
		}
	}

	return
}

func test5() (err error) {

	defer func() {
		fmt.Println("defer err5:", err)
	}()

	if _, err = os.Open("xxx"); err == nil {
		if err = json.Unmarshal([]byte("{}"), &struct{}{}); err == nil {
			fmt.Println("OK")
		}
	}

	return
}
