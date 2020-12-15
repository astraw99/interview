package main

import (
	"fmt"
	"runtime"
)

func main() {
	Foo()
}
func Foo() {
	fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println(file, line, ok)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, file, line, ok := runtime.Caller(2)
	fmt.Println(file, line, ok)
	return runtime.FuncForPC(pc).Name()
}
