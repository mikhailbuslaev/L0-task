package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)
// Continuous writing to data channel
func write(data chan int, quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("Publisher quit")
			return
		default:
			time.Sleep(500*time.Millisecond)
			data <- rand.Int()
		}
	}
}
// Listennig data channel
func listen(data chan int, quit chan bool) {
	for {
		select {
		case <- quit:
			fmt.Println("Listener quit")
			return
		case <- data:
			fmt.Fprintf(os.Stdout, "Listener receive message '%d'.\n", <-data)
		}
	}
}
// handle safety quit for writer and listener
func handleQuit(quit chan bool, timeout time.Duration) {
	time.Sleep(timeout)
	for {
		//endless loop is safety here, because it finish after finish of writer and listener
		quit <- true
	}
} 

func main() {
	data := make(chan int)
	quit := make(chan bool)
	var duration string
	fmt.Print("Enter duration of work: ")
	fmt.Scan(&duration)//cmd scan
	timeout, _ := time.ParseDuration(duration)

	go handleQuit(quit, timeout)
	go write(data, quit)
	listen(data, quit)
}