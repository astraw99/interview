package main

import "fmt"

func main() {
	PaseStudent()
}

type student struct {
	Name string
	Age  int
}

func PaseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	// 都指向了同一个stu的内存指针
	fmt.Println(m)
}
