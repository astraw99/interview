package main

import "fmt"

func main() {
	ParseStudent()
}

type student struct {
	Name string
	Age  int
}

func ParseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhang", Age: 22},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 24},
	}
	for _, stu := range stus {
		// 都指向了同一个stu的内存指针，为什么？
		// 因为 for range 中的 v 只会声明初始化一次
		// 不会每次循环都初始化，最后赋值会覆盖前面的
		fmt.Printf("%p\n", &stu)

		// 1. bad
		//m[stu.Name] = &stu

		// 2. good
		newStu := stu
		m[stu.Name] = &newStu
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
