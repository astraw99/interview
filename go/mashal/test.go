package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := ""
	jsonOut := ""

	err := json.Unmarshal([]byte(str), &jsonOut)

	// err: unexpected end of JSON input
	fmt.Println(err)
}
