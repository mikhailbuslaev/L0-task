package main

import (
	"strings"
	"fmt"
)

func check(input string) bool{
	groups := make(map[string]string)
	input = strings.ToLower(input)
	arr := strings.Split(input, "")
	for i := range arr {
		if groups[arr[i]] == "" {
			groups[arr[i]] = arr[i]
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(check("AbCdgheF"))
}