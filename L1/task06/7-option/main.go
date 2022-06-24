package main

import (
	"time"
)

// we can stop goroutine via select without 
//special channel and without close() if transmit in our data channel special quit message
func doSmthng(c chan string) {
	for {
		select {
		case message:= <-c:
			if message =="quit" {
				println("goroutine exit...")
				return
			} else {
				println("goroutine do smthng...")
				time.Sleep(500*time.Millisecond)
			}
		}
	}
}

func main() {
	c := make(chan string)
	go doSmthng(c)
	for i:=0;i<5;i++ {
		c <- "message example"
	}
	c <- "quit"
	time.Sleep(1*time.Second)
}