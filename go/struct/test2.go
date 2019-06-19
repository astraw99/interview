package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
	Age  int32
}

type Test2 struct {
	Name string
	Age  int32
}

func main() {
	t1 := Test{Name: "test", Age: 20}
	t2 := Test(Test2{Name: "test", Age: 20})
	//fmt.Println(t1 == t2)

	fmt.Println(reflect.TypeOf(t1))
	fmt.Println(reflect.TypeOf(t2))
	fmt.Println(reflect.TypeOf(t1) == reflect.TypeOf(t2))

	fmt.Println(reflect.ValueOf(t1))
	fmt.Println(reflect.ValueOf(t2))
	fmt.Println(reflect.ValueOf(t1.Name) == reflect.ValueOf(t2.Name))
}
