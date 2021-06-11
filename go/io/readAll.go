package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*absPath, err := filepath.Abs("go/iota/test.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(absPath)

	fileInfo, err := os.Open(absPath)*/

	fileInfo, err := os.Open("go/json/scientificNum.go")
	if err != nil {
		panic(err)
	}

	contentBytes, err := io.ReadAll(fileInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(contentBytes))
}
