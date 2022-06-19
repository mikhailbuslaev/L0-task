package main

import (
	"fmt"
)

func main() {
	groups := make(map[int][]float32)
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for i := range temp {
		group := 10 * int(temp[i]/10)
		if groups[group] == nil {
			groups[group] = make([]float32, 0, 10)
			groups[group] = append(groups[group], temp[i])
		} else {
			groups[group] = append(groups[group], temp[i])
		}
	}
	fmt.Println(groups)
}
