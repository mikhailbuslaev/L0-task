package main

import (
	"time"
)

func doSmthng() {
	fmt.Println("Goroutine running...")
	time.Sleep(30*time.Second)
}

//1 option: close main goroutine without blocking
func main() {
 	go doSmthng()
}