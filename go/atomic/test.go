package main

import (
	"fmt"
	"sync/atomic"
)

func main()  {
	var a int32

	atomic.StoreInt32(&a, 100)
	fmt.Println(atomic.LoadInt32(&a))

	atomic.AddInt32(&a, 50)
	fmt.Println(atomic.LoadInt32(&a))

	atomic.SwapInt32(&a, 200)
	fmt.Println(atomic.LoadInt32(&a))

	atomic.CompareAndSwapInt32(&a, 200, 300)
	fmt.Println(atomic.LoadInt32(&a))
}
