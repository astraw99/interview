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

	/*for v := range t.C {
		fmt.Println(v)
	}*/

	//t.Stop()
}
