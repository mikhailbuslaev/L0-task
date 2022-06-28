package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	reverseOption *bool
	columnNumber  *int
	numbersOption *bool
	uniqueOption  *bool
)

func parseRows(fileName string) ([]string, error) {
	// open file
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return []string{""}, err
	}
	// split buffer by newline
	rows := strings.Split(strings.ReplaceAll(string(buf), "\r\n", "\n"), "\n")
	// if columnNumber flag set, split rows by [space] symbol and make 2d-array
	if *columnNumber != -1 { // -1 is default value, because we can`t set it as 0
		array2d := make(map[int][]string)
		for i := range rows {
			array2d[i] = make([]string, 0, 10)                              // create columns array
			array2d[i] = append(array2d[i], strings.Split(rows[i], " ")...) // append to columns array
		}
		rows = make([]string, 0, 10) // recreate rows array
		for i := range array2d {
			if len(array2d[i]) > *columnNumber {
				rows = append(rows, array2d[i][*columnNumber]) // append columns to output
			}
		}
	}

	if *uniqueOption {
		karta := make(map[string]string)
		for _, v := range rows {
			if len(karta[v]) == 0 {
				karta[v] = v
			}
		}
		rows = make([]string, 0, 10)
		for _, v := range karta {
			rows = append(rows, v)
		}
	}

	return rows, nil
}

func printRows(rows []string) {
	length := len(rows)
	if *reverseOption { // checking for reverse flag
		for i := length - 1; i > -1; i-- { // reverse printing
			fmt.Println(rows[i])
		}
	} else {
		for i := 0; i < length; i++ { // direct printing
			fmt.Println(rows[i])
		}
	}
}

func sortRows(rows []string) []string {
	if *numbersOption { // flag for numeric or string sorting
		nums := make([]int, 0, 10) // ints sort
		for i := range rows {
			if v, err := strconv.Atoi(rows[i]); err == nil {
				nums = append(nums, v)
			}
		}
		sort.Ints(nums)
		rows = make([]string, 0, 10)
		for i := range nums {
			rows = append(rows, strconv.Itoa(nums[i]))
		}
		return rows
	}
	sort.Strings(rows) // strings sort
	return rows
}

func main() {
	sortCommand := flag.NewFlagSet("dev03_sort", flag.ExitOnError)
	reverseOption = sortCommand.Bool("r", false, "reverse sorting order")
	columnNumber = sortCommand.Int("k", -1, "number of sorted column")
	numbersOption = sortCommand.Bool("n", false, "number sorting by ascendent")
	uniqueOption = sortCommand.Bool("u", false, "print only unique records")

	if len(os.Args) < 2 {
		fmt.Println("expected 'dev03_sort' subcommand")
		os.Exit(1)
	}

	if os.Args[1] == "dev03_sort" {
		sortCommand.Parse(os.Args[2:])
		rows, err := parseRows(sortCommand.Arg(0))
		if err != nil {
			fmt.Println("cannot parse file '" + sortCommand.Arg(0) + "'")
			os.Exit(1)
		}
		rows = sortRows(rows)
		printRows(rows)
	}
}
