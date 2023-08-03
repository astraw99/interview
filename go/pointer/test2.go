package main

import "fmt"

type pvc struct {
	Name             string
	StorageClassName *string
}

func main() {
	sc := "test-sc"
	pvcs := []*pvc{&pvc{
		Name:             "test-name",
		StorageClassName: &sc,
	}}

	scTarget := "test-sc"
	for _, pvc := range pvcs {
		if pvc.StorageClassName == &scTarget {
			fmt.Println("sc match ok 1", pvc.StorageClassName, &scTarget)
		} else {
			fmt.Println("sc not match 1", pvc.StorageClassName, &scTarget) // hit here: sc not match 1 0xc000010230 0xc000010240
		}
	}

	for _, pvc := range pvcs {
		if *pvc.StorageClassName == scTarget {
			fmt.Println("sc match ok 2", *pvc.StorageClassName, scTarget) // hit here: sc match ok 2 test-sc test-sc
		} else {
			fmt.Println("sc not match 2", *pvc.StorageClassName, scTarget)
		}
	}
}
