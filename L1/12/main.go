package main

import (
	"fmt"
)

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	group := make(map[string]string)
	for i := range input {
		if group[input[i]] == "" {
			group[input[i]] = input[i]
		}
	}
	fmt.Println(group)
}
