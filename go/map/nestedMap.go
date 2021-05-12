package main

import "fmt"

func main()  {
	// nested map
	outerMap := make(map[string]map[string][]string)
	outerMap["a"] = make(map[string][]string) // innerMap: will panic if not init!!
	outerMap["b"] = make(map[string][]string) // innerMap: will panic if not init!!

	outerMap["a"]["a1"] = []string{
		"test1",
	}
	outerMap["a"]["a2"] = []string{
		"test1",
		"test2",
	}
	outerMap["b"]["b1"] = []string{
		"test1",
	}
	outerMap["b"]["b2"] = []string{
		"test1",
		"test2",
	}

	fmt.Println(outerMap)
}
