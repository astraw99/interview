package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1000000)) // duration
	//ctx, cancel := context.WithTimeout(context.Background(), 0) // 0 => ctx.Done immediately
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	chiHanBao(ctx)
	defer cancel()
}

func chiHanBao(ctx context.Context) {
	n := 0
	rand.Seed(time.Now().Unix())
	for {
		select {
		// timeout 或 cancel() 都会 -> ctx.Done
		case <-ctx.Done():
			fmt.Println("stop")
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
