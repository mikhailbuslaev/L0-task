package main

import (
    "fmt"
    "time"
	"os"
)

func listen(data chan int) {
	fmt.Println("Listening goroutine running...")
	for {
		val, ok := <- data
		if !ok {
			fmt.Println("Listening goroutine closing...")
			return
		}
		fmt.Fprintf(os.Stdout, "Received new message: '%d'\n", val)
		time.Sleep(400*time.Millisecond)
	}
}

func publish(data chan int) {
	fmt.Println("Publishing goroutine running...")
	for i := 0; i<5; i++{
		data <- i
	}
	fmt.Println("Init goroutines close...")
	close(data)
	fmt.Println("Publish goroutine closing...")
}
// If you already have used channels, you can stop goroutine via close(your channel) outside gourutine
func main() {
	data := make(chan int)
	go listen(data)
	publish(data)
	time.Sleep(1*time.Second)
}