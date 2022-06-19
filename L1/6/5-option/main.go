package main

import (
	"fmt"
	"sync"
	"time"
)

func doSmthng() {
	fmt.Println("Goroutine do something ...")
	time.Sleep(400 * time.Millisecond)
}

// 5 option, stop goroutine via waitgroup
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine running...")
		doSmthng()
		defer wg.Done()
	}(&wg)
	wg.Wait()
}
