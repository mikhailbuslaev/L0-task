package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "главрыба"
	var output string = ""
	array := strings.Split(input, "")
	mid := int(float64(len(array)) / 2)
	length := len(array)
	for i := 0; i < mid; i++ {
		array[i], array[length-i-1] = array[length-i-1], array[i]// a,b = b, a until half of array
	}
	for i := range array {
		output += array[i]
	}
	fmt.Println(output)
}
