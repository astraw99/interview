package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var num int

func main() {
	m = new(sync.RWMutex)
	num = 100

	// 多个同时读写，被读锁占用时，写锁等待
	go read(1)
	go read(2)
	go write(101)
	go write(102)
	go read(3)

	time.Sleep(5 * time.Second)
}

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading before sleep", num)
	time.Sleep(2 * time.Second)
	println(i, "reading after sleep", num)
	m.RUnlock()

	println(i, "read over")
}

func write(i int) {
	println("writing before", num)

	m.Lock()
	defer m.Unlock()

	num = i
	println("writing after", num)
}
