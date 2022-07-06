package main

import (
	"strings"
	"sort"
	"fmt"
)

func findAnagrams(arrP *[]string) map[string]*[]string{
	arr := *arrP
	anagrams := make(map[string]*[]string)// create map with groups of anagrams
	for i:= range arr {
		word := strings.Split(arr[i], "")// split input string into letters
		sort.Strings(word)// sort letters
		key := strings.Join(word, "")// join letters
		if anagrams[key] == nil {// check if group does not exist
			anagrams[key] = new([]string)
			*anagrams[key] = append(*anagrams[key], arr[i])
		} else {
			*anagrams[key] = append(*anagrams[key], arr[i])
		}
	}
	for i := range anagrams {// loop over angrams group
		if len(*anagrams[i]) == 1 {
			delete(anagrams, i)// reject single-worded groups
		} else {
			sort.Strings(*anagrams[i])// sort group
			fmt.Println(*anagrams[i])
		}
	}
	return anagrams
}

func main() {
	arr := &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(findAnagrams(arr))
}