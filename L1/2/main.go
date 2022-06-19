package main

func countSqr(input int) int{
	return input*input
}

func main() {
	array := [5]int{2,4,6,8,10}
	c := make(chan int)
	for _, v := range array {
		go func(c chan int, v int) {
			c <- countSqr(v)
		}(c, v)
	}

	for range array {
		println(<-c)
	}
}