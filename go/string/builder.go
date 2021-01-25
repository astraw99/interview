package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder1 strings.Builder

	builder1.WriteString("Hello")
	fmt.Printf("%q, %d, %d\n", builder1.String(), builder1.Len(), builder1.Cap())

	// 扩容为原 len*2 + n
	builder1.Grow(10)
	fmt.Printf("%q, %d, %d\n", builder1.String(), builder1.Len(), builder1.Cap())

	// 不能被 copy
	var b2 = builder1
	b2.Grow(100) // panic

	builder1.Reset()
	fmt.Printf("%q, %d, %d\n", builder1.String(), builder1.Len(), builder1.Cap())
}
