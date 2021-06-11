package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	unmarshal() // 会转为科学计数法: 8.589934592e+09
	decode()    // ok
}

func unmarshal() {
	body := `{"message": "success", "data": 8589934592}`
	v := make(map[string]interface{})

	// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
	// float64, for JSON numbers
	err := json.Unmarshal([]byte(body), &v)
	fmt.Println(v, err)
}

func decode() {
	body := `{"message": "success", "data": 8589934592}`
	v := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewReader([]byte(body)))
	d.UseNumber()
	err := d.Decode(&v)
	fmt.Println(v, err)
}
