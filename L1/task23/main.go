package main

import "fmt"

func cut(array []int, i int) []int {
	return append(array[:i], array[i+1:]...)
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(cut(input, 2))
}
