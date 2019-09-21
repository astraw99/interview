package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]int{"key": 1}

	var b bytes.Buffer
	// json.encode 会默认加换行符 \n
	json.NewEncoder(&b).Encode(data)

	raw, _ := json.Marshal(data)

	if b.String() == string(raw) {
		fmt.Println("same encoded data")
	} else {
		fmt.Printf("'%s' != '%s'", raw, b.String())
		//prints:
		//'{"key":1}' != '{"key":1}\n'
	}
}
