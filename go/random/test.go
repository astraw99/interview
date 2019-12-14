package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 每次运行重置随机值种子

	fmt.Println(rand.Intn(2)) // [0, 2) 左闭右开随机值
}
