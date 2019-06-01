package main

import (
	"fmt"
	"sync"
)

func main() {
	ua := UserAges{}
	ua.Add("wang", 20)

	fmt.Println(ua.Get("wang"))
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	// map 一定要先make初始化才能赋值
	ua.ages = make(map[string]int)
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
