package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex// доступ к чтению Счетчика будет заблокирован 
	//первой горутиной, которая захочет его прочитать
	Value int
}

func main() {
	c := &Counter{Value: 0}
	wg := sync.WaitGroup{}
	wg.Add(5)// для завершения работы 5 сумматоров
	for i := 0; i < 5; i++ {
		go func(c *Counter, i int) {// запускаем сумматоры
			fmt.Fprintf(os.Stdout, "Summator №%d goes...\n", i)
			for j := 0; j < 5; j++ {
				time.Sleep(1 * time.Second)
				fmt.Fprintf(os.Stdout, "Summator №%d increment...\n", i)
				c.Lock()// закрываем доступ другим, другие будут ждать разблокировки
				c.Value++//5 раз суммируем
				c.Unlock()// открываем доступ
			}
			wg.Done()// завершаем работу сумматора
		}(c, i)
	}
	wg.Wait()
	fmt.Println(c.Value)
}
