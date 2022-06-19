package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	Value int
}

func main() {
	c := &Counter{Value: 0}
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(c *Counter, i int) {
			fmt.Fprintf(os.Stdout, "Summator №%d goes...\n", i)
			for j := 0; j < 5; j++ {
				time.Sleep(1 * time.Second)
				fmt.Fprintf(os.Stdout, "Summator №%d increment...\n", i)
				c.Lock()
				c.Value++
				c.Unlock()
			}
			wg.Done()
		}(c, i)
	}
	wg.Wait()
	fmt.Println(c.Value)
}
