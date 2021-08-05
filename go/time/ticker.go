package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)

	select {
	case v := <-t.C:
		fmt.Println(v)
	}

	t.Stop() // just turn off ticker, not close channel!!

	if len(t.C) == 0 {
		fmt.Println("len0", t.C)
	}

	for v := range t.C {
		fmt.Println(v) // deadlock: will not stop forever
	}
}
