package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Println("sepIndex:", sepIndex)

	dir1 := path[:sepIndex:sepIndex] // full slice expression: [low : high : max] len = high - low, cap = max - low
	fmt.Println("dir1:", string(dir1), len(dir1), cap(dir1))

	dir2 := path[sepIndex+1:]
	fmt.Println("dir2:", string(dir2), len(dir2), cap(dir2))

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

	fmt.Println("new path =>", string(path))


	// =============================================

	test := []string{"a", "b", "c", "d", "e"}
	index := 4

	test = append(test[:index], test[index + 1:]...)

	// boundsSliceB    // s[x:y], 0 <= x <= y failed (but boundsSliceA didn't happen)
	//fmt.Println(test[index + 1:]) // will panic?

	fmt.Println(test)
}
