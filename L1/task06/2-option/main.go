package main

import (
	"time"
	"fmt"
)

func doSmthng(quit chan bool) {
	fmt.Println("Goroutine running...")
	for {
		select {
		case <- quit:
			fmt.Println("Goroutine closing...")
			return
		}
	}
}

// 2 option: close goroutine via select{quit channel}
func main() {
	quit := make(chan bool)
 	go doSmthng(quit)
	
	time.Sleep(2*time.Second)
	quit <- true
	time.Sleep(1*time.Second)
}