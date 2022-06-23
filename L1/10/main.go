package main

import (
	"fmt"
)

func main() {
	groups := make(map[int][]float32)
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}// array of temperatures
	for i := range temp {
		group := 10 * int(temp[i]/10)// we can divide into groups by this way, because int() rounds to smaller value
		if groups[group] == nil {
			groups[group] = make([]float32, 0, 10)// create group in map if doesnt exists
			groups[group] = append(groups[group], temp[i])// append first element
		} else {
			groups[group] = append(groups[group], temp[i])//append next element to group
		}
	}
	fmt.Println(groups)
}
