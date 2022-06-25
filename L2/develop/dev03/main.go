package main

import (
	"flag"
	"strings"
	"os"
	"sort"
	"fmt"
)

func parseRows(fileName string) ([]string, error) {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return []string{""}, err
	}
	return strings.Split(string(buf), "\n"), nil
}

func sortRows(rows []string) []string{
	sort.Strings(rows)
	return rows
}

func printRows(rows []string) {
	for _, v := range rows {
		fmt.Println(v)
	}
}

func main() {
	sortCommand := flag.NewFlagSet("dev03_sort", flag.ExitOnError)
	if len(os.Args) < 2 {
        fmt.Println("expected 'dev03_sort' subcommand")
        os.Exit(1)
    }

	if os.Args[1] == "dev03_sort" {
		sortCommand.Parse(os.Args[2:])
        fmt.Println("subcommand 'dev03_sort'")
        fmt.Println("  tail:", sortCommand.Args())
		rows, err := parseRows(sortCommand.Arg(0))
		if err != nil {
			fmt.Println("cannot parse file '"+sortCommand.Arg(0)+"'")
        	os.Exit(1)
		}

		rows = sortRows(rows)
		printRows(rows)
	}
}