package main

import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
)

func isInt(s string) bool{
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func unwrapString(input string) (string, error){
	arr := strings.Split(input, "")
	buf := bytes.Buffer{}
	newArr := make([]string, 0, len(arr))
	length := len(arr)
	for i := 0; i < length-1; i++ {
		switch {
		case isInt(arr[i]) && isInt(arr[i+1]):
			return "", fmt.Errorf("incorrect sequence: '"+arr[i]+arr[i+1]+"'")

		case !isInt(arr[i]) && isInt(arr[i+1]):
			num, _ := strconv.Atoi(arr[i+1])
			for j := 0; j < num; j++ {
				newArr = append(newArr, arr[i])
			}

		case isInt(arr[i]) && !isInt(arr[i+1]):// do nothing
		default:
			newArr = append(newArr, arr[i])
		}
	}
	newArr = append(newArr, arr[len(arr)-1])
	// Assembly array of strings in one
	return strings.Join(newArr, ""), nil
}

func main() {
	fmt.Println(unwrapString("a4bc2d5e"))
	fmt.Println(unwrapString(string([]rune("a4фвывф5e"))))
}