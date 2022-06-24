package main

import (
	"runtime"
	"time"
)

func exampleFunc(quit chan bool) {
	for {
		select {
		case <-quit:
			println("Bye")
			runtime.Goexit()//we can use Goexit() instead of return
		default:
			println("Do smthng...")
			time.Sleep(1*time.Second)
		}
	}
}

func main() {
	quit := make(chan bool)
	go exampleFunc(quit)
	time.Sleep(5*time.Second)
	quit <- true
	time.Sleep(1*time.Second)
}