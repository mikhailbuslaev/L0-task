package main

func countSqr(input int) int{
	return input*input
}

func main() {
	c := make(chan int)
	array := [5]int{2,4,6,8,10}
	sum := 0
	for _, v := range array {
		go func(c chan int, v int) {
			c <- countSqr(v)
		}(c, v)
	}
	
	for range array {
		sum += <- c
	}
	print(sum)
}