package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	s := make([]string, 5, 10)
	modifySlice(s)
	fmt.Println(s)// [3 2 3]
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
