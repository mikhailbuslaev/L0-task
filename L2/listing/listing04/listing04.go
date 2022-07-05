package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i// 10 раз пишем
		}
	}()
	for n := range ch {// читаем здесь
		println(n)
	}
}
// дедлок на 11 чтении