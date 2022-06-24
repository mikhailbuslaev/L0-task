package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input := make(chan int)
	counted := make(chan int)
	// Int generator
	go func() {
		i := 0
		for {
			time.Sleep(400 * time.Millisecond)
			i++
			input <- i
		}
	}()
	// Counter
	go func() {
		for {
			counted <- 2 * <-input
		}
	}()
	// Printer
	func() {
		for {
			fmt.Fprintf(os.Stdout, "Received value: %d \n", <-counted)
		}
	}()
}
