package main

import (
	"strings"
	"sort"
	"fmt"
)

func findAnagrams(arrP *[]string) map[string]*[]string{
	arr := *arrP
	anagrams := make(map[string]*[]string)
	for i:= range arr {
		word := strings.Split(arr[i], "")
		sort.Strings(word)
		key := strings.Join(word, "")
		if anagrams[key] == nil {
			anagrams[key] = new([]string)
			*anagrams[key] = append(*anagrams[key], arr[i])
		} else {
			*anagrams[key] = append(*anagrams[key], arr[i])
		}
	}
	for i := range anagrams {
		if len(*anagrams[i]) == 1 {
			delete(anagrams, i)
		} else {
			sort.Strings(*anagrams[i])
			fmt.Println(*anagrams[i])
		}
	}
	return anagrams
}

func main() {
	arr := &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(findAnagrams(arr))
}