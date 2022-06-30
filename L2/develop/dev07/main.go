package main

import (
	"fmt"
	"time"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan bool {
	done := make(chan bool)
	for i := range channels {
		go func(channel <-chan interface{}, done chan bool) {
			_, ok := <-channel
			if !ok {
				fmt.Println("catched closed channel!")
				done <- true
			}
		}(channels[i], done)
	}
	return done
}

func main() {
	<-or(
		sig(5*time.Second),
		sig(9*time.Second),
		sig(15*time.Second),
	)
	fmt.Println("done channel readed...")
}
