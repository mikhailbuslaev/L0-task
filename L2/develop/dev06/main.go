package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	fieldNumLeft  *int
	fieldNumRight *int
	delimiter     *string
	onlyDelimited *bool
)

func cut(s string) []string {
	leftBorder := *fieldNumLeft
	rightBorder := *fieldNumRight
	strArray := strings.Split(s, *delimiter)
	if len(strArray) <= 1 && *onlyDelimited {
		return []string{""}
	}
	switch {
	case leftBorder > -1 && rightBorder > -1:
		cutArray := make([]string, rightBorder-leftBorder+1)
		copy(cutArray, strArray[leftBorder:rightBorder+1])
		return cutArray
	case leftBorder > -1 && rightBorder == -1:
		cutArray := make([]string, len(strArray)-leftBorder)
		copy(cutArray, strArray[leftBorder:])
		return cutArray
	case leftBorder == -1 && rightBorder > -1:
		cutArray := make([]string, rightBorder+1)
		copy(cutArray, strArray[:rightBorder+1])
		return cutArray
	}
	return strArray
}

func main() {
	cutCommand := flag.NewFlagSet("dev06_cut", flag.ExitOnError)
	fieldNumLeft = cutCommand.Int("fl", -1, "number of sorted column")
	fieldNumRight = cutCommand.Int("fr", -1, "number of sorted column")
	delimiter = cutCommand.String("d", " ", "number of sorted column")
	onlyDelimited = cutCommand.Bool("s", false, "something")

	if len(os.Args) < 2 {
		fmt.Println("expected 'dev06_cut' subcommand")
		os.Exit(1)
	}

	if os.Args[1] == "dev06_cut" {
		cutCommand.Parse(os.Args[2:])
		cutStr := cut(cutCommand.Arg(0))
		for _, v := range cutStr {
			fmt.Println(v)
		}
	}
}
