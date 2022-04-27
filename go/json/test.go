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

type Project struct {
	Key   string `json:"key,omitempty"` // del empty val
	Value string `json:"-"`             // ignore field
	Name  string `json:"-,"`            // name to "-"
	Num   int64  `json:"num,string"`    // convert to string: "num":"1"
}

type JiraHttpReqField struct {
	Project     `json:",inline"` // inline to flatten field
	Summary     string           `json:"summary"`
	Description string           `json:"description"`
}

func main() {
	str := "{\"header\":{\"code\":200,\"msg\":\"success\",\"desc\":\"请求成功\",\"debug\":{\"cost_time\":\"51.622\"}},\"body\":{\"MIIN0002\":{\"14138\":0,\"16279\":0},\"MIIN0003\":{\"14138\":0,\"16279\":0}}}"
	test := Test{}
	json.Unmarshal([]byte(str), &test)

	fmt.Println(test)

	bytes, _ := json.Marshal(test)
	fmt.Println(string(bytes))

	dataProject := Project{
		Key:   "",
		Value: "val",
	}
	dataJiraHttpReqField := &JiraHttpReqField{
		Project:     dataProject,
		Summary:     "Summary",
		Description: "Description",
	}
	data, _ := json.Marshal(dataJiraHttpReqField)
	fmt.Println(string(data))

	testF()
}

func testF() {
	str := []byte(`{"command":["echo hello","echo world"]}`)
	payload := SubProcessPayload{}
	err := json.Unmarshal(str, &payload)
	fmt.Println(payload, err)

	payload2 := SubProcessPayload{
		Command: []string{"echo hello", "echo world"},
	}
	data, err := json.Marshal(&payload2)

	fmt.Println(string(data), err)
}

type SubProcessPayload struct {
	Command []string `json:"command"`
}
