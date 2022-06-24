package main

import (
	"fmt"
	"math"
	"strings"
)

func replace(input, num int64, value bool) int64 {
	// Parse to bits array
	bits := strings.Split(fmt.Sprintf("%064b", input), "") // Cheat way
	// Replace
	if !value {
		bits[num] = "0"
	}
	if value {
		bits[num] = "1"
	}
	// Convert to int64
	length := len(bits)
	out := 0.00
	for i := length - 1; i >= 0; i-- {
		if bits[i] == "1" {
			out += math.Pow(2, float64(length-i-1))
		}
	}
	return int64(out)
}

func main() {
	fmt.Println(replace(1111, 2, false))
	fmt.Println(replace(1111, 2, true))
}
