package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["aaa"] = "AAA"
	m["bbb"] = "BBB"
	m["ccc"] = "CCC"
	m["ddd"] = "DDD"
	m["eee"] = "EEE"

	// range map 为随机序输出
	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println("==================")

	s := []string{"aaa", "bbb", "ccc", "ddd", "eee"}

	// range slice 为索引序输出
	for i, v := range s {
		fmt.Println(i, v)
	}
}
