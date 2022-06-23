package main

import (
	"fmt"
)

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	group := make(map[string]string)
	for i := range input {
		if group[input[i]] == "" {// add to map if it first example
			group[input[i]] = input[i]
		}
	}
	fmt.Println(group)
}
