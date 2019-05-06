package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	for i := 0; i < 20; i++ {
		go func() {
			for {
				b := make([]byte, 10)
				os.Stdin.Read(b) // will block
			}
		}()
	}

	fmt.Println(runtime.NumGoroutine())

	select {}
}
