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

	// 1. uint32 减法
	var b uint32 = 10
	delta := int32(-3)
	//atomic.AddUint32(&b, uint32(int32(-3))) // constant overflows uint32
	atomic.AddUint32(&b, uint32(delta))
	fmt.Println(atomic.LoadUint32(&b))

	// 2. uint32 减法
	n := -3
	fmt.Println(atomic.AddUint32(&b, ^uint32(-n - 1)))
}
