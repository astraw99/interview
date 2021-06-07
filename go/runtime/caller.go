package main

import (
	"fmt"
	"runtime"
	"strings"
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
	callers()
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

func callers() {
	pcs := make([]uintptr, 2)
	_ = runtime.Callers(0, pcs)

	var loggerPackage string
	fmt.Println(runtime.Version())

	if runtime.Version() >= "go1.14" {
		loggerPackage = getPackageName(runtime.FuncForPC(pcs[0]).Name())
	} else {
		loggerPackage = getPackageName(runtime.FuncForPC(pcs[1]).Name())
	}

	fmt.Println("loggerPackage:", loggerPackage)
}

// getPackageName to get the package name
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")

		// [abc/def].main
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
