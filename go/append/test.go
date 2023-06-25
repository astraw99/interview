package main

import "fmt"

type pod struct {
	name      string
	namespace string
}

func main() {
	var s []interface{}

	s = append(s, pod{
		name:      "pod-1",
		namespace: "default",
	})

	s = append(s, "just a string")

	var pods []pod
	for _, o := range s {
		//pods = append(pods, o.(pod)) // will panic
		p, ok := o.(pod)
		if ok {
			pods = append(pods, p)
		}
	}

	fmt.Println(pods)
}
