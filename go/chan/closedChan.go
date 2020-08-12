package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func main() {
	TestReadFromClosedChan(&testing.T{})
}

func TestReadFromClosedChan(t *testing.T) {
	asChan := func(vs ...int) <-chan int {
		c := make(chan int)
		go func() {
			for _, v := range vs {
				c <- v
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
			close(c)
		}()
		return c
	}
	merge := func(a, b <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			for {
				select {
				case v := <-a:
					fmt.Println("a:", v)
					c <- v
				case v := <-b:
					fmt.Println("b:", v)
					c <- v
				}
			}
		}()
		return c
	}

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)

	fmt.Println(len(a), len(b))

	c := merge(a, b)
	fmt.Printf("%T", c)
	//os.Exit(0)
	for v := range c {
		fmt.Println(v)
	}

	// 为什么会一直输出0？
}
