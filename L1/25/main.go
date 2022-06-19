package main

import (
	"fmt"
)

func sleep(ticks int) {
	fmt.Println("sleep run...")
	c := make(chan bool)
	// Timer
	go func(c chan bool) {

		for i:=0;true;i++ {
			if i%100000000 == 0 {
				c <- true
			}
		}
	}(c)
	// Counter
	func(c chan bool, ticks int) {
		for i:=0;i<ticks;i++ {
			<-c
		}
	}(c, ticks)

	fmt.Println("sleep done")
}

func main() {
	sleep(100)
}
