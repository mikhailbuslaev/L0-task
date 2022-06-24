package main

import (
	"fmt"
)

func sortByGroup(array []int) map[int][]int {// copied from 10 task, little changes

	groups := make(map[int][]int)
	for i := range array {
		group := 10 * int(float64(array[i])/10)
		if groups[group] == nil {
			groups[group] = make([]int, 0, 10)
			groups[group] = append(groups[group], array[i])
		} else {
			groups[group] = append(groups[group], array[i])
		}
	}
	return groups
}

func main() {
	array1 := []int{1, 2, 3, 7, 95, 63, 22, 48, 49, 66, 98}
	array2 := []int{9, 82, 51, 47, 39, 76, 5, 93, 55, 3, 63}
	intersection := make([]int, 0, 10)
	group1 := sortByGroup(array1)// i want divide array into array of groups for better speed
	group2 := sortByGroup(array2)// if compared group doesnt exists, i just quit it and go to next group
								// it allows me except "hopeless" groups
								// example: first map: map{"10":{11,12,18},"0":{1,2,3}}
								// second map: map{"10":{17,14,11}, "20":{21,22,23}}
								// if we compare first map with second, at the start of comparing 
								// we just drop "20" group, because in first map we dont have it
	for i := range group1 {
		if group2[i] != nil {
			length1 := len(group1[i])
			length2 := len(group2[i])
			for j := 0; j < length1; j++ {
				for k := 0; k < length2; k++ {
					if group1[i][j] == group2[i][k] {
						intersection = append(intersection, group1[i][j])
					}
				}
			}
		}
	}
	fmt.Println(intersection)
}
