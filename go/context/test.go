package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	chiHanBao(ctx)
	defer cancel()
}

func chiHanBao(ctx context.Context) {
	n := 0
	rand.Seed(time.Now().Unix())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop \n")
			return
		default:
			incr := rand.Intn(5)
			fmt.Println("rand", incr)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}
