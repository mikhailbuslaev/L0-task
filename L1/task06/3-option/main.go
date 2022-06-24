package main

import (
	"context"
	"time"
	"fmt"
)

func doSmthng(ctx context.Context) {
	fmt.Println("Goroutine running...")
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Goroutine stopping...")
			return
		default:
			fmt.Println("Goroutine do something ...")
			time.Sleep(400*time.Millisecond)
		}
	}
}
// 3 option, stop goroutine via context
func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	go doSmthng(ctx)
	time.Sleep(2*time.Second)
	cancelCtx()
	time.Sleep(2*time.Second)
}