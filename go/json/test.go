package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Header struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Desc string `json:"desc"`
	} `json:"header"`
	Body map[string]map[string]int `json:"body"`
}

func main() {
	str := "{\"header\":{\"code\":200,\"msg\":\"success\",\"desc\":\"请求成功\",\"debug\":{\"cost_time\":\"51.622\"}},\"body\":{\"MIIN0002\":{\"14138\":0,\"16279\":0},\"MIIN0003\":{\"14138\":0,\"16279\":0}}}"
	test := Test{}
	json.Unmarshal([]byte(str), &test)

	fmt.Println(test)

	bytes, _ := json.Marshal(test)
	fmt.Println(string(bytes))
}
