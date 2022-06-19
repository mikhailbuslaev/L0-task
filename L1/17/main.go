package main

import "fmt"

func binarySearch(a []int, search int) int {
	mid := len(a) / 2
	var result int
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	return result
}

func main() {
	arr := []int{1, 5, 8, 9, 16, 56, 77, 89, 111, 150}
	result := binarySearch(arr, 77)
	fmt.Println(result)
}
