package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

	start, err := os.Stat(os.Args[1])
	if err != nil || !start.IsDir() {
		os.Exit(-1)
	}

	var targets []string
	_ = filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fpath)
		return nil
	})

	for _, target := range targets {
		func() {
			f, err := os.Open(target)
			if err != nil {
				fmt.Println("bad target:", target, "error:", err)
				return // 与 break 等价
			}
			defer f.Close() //ok
			//do something with the file...
		}() // 匿名函数结束，会执行 defer
	}
}
