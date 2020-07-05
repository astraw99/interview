package main

import (
	"fmt"
	"sync"
)

type myLocker sync.Locker

func main() {
	var locker myLocker
	locker.Lock()
	a := 100
	locker.Unlock()

	fmt.Println(a)
}
