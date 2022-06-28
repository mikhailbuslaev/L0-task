package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mikhailbuslaev/wb-tasks/l2/dev05/greper"
	"github.com/mikhailbuslaev/wb-tasks/l2/dev05/parser"
	"github.com/mikhailbuslaev/wb-tasks/l2/dev05/printer"
)

var (
	afterLength    *int
	beforeLength   *int
	contextLength  *int
	countOption    *bool
	ignoreRegister *bool
	invertOption   *bool
	fixedOption    *bool
	lineNumOption  *bool
)

type App struct {
	Parser  *parser.Parser
	Greper  greper.Greper
	Printer *printer.Printer
}

func (a *App) Set() {
	// greper set
	switch {
	case *fixedOption:
		a.Greper = &greper.FixedGreper{}
	case *invertOption:
		a.Greper = &greper.InvertGreper{}
	default:
		a.Greper = &greper.DefaultGreper{Params: greper.Params{*beforeLength, *afterLength, *lineNumOption}}
	}

	// parser set
	a.Parser = &parser.Parser{parser.Params{*ignoreRegister}}

	// printer set
	a.Printer = &printer.Printer{printer.Params{*countOption}}
}

func main() {
	grepCommand := flag.NewFlagSet("dev05_grep", flag.ExitOnError)
	afterLength = grepCommand.Int("a", 0, "number of sorted column")
	beforeLength = grepCommand.Int("b", 0, "number of sorted column")
	contextLength = grepCommand.Int("c", 0, "number of sorted column")
	countOption = grepCommand.Bool("count", false, "number of sorted column")
	ignoreRegister = grepCommand.Bool("i", false, "number sorting by ascendent")
	invertOption = grepCommand.Bool("v", false, "number sorting by ascendent")
	fixedOption = grepCommand.Bool("F", false, "number sorting by ascendent")
	lineNumOption = grepCommand.Bool("n", false, "number of sorted column")

	if len(os.Args) < 2 {
		fmt.Println("expected 'dev05_grep' subcommand")
		os.Exit(1)
	}

	if os.Args[1] == "dev05_grep" {
		grepCommand.Parse(os.Args[2:])
		a := &App{}
		*afterLength = *afterLength + *contextLength
		*beforeLength = *beforeLength + *contextLength
		a.Set()
		rows, pattern, err := a.Parser.Parse(grepCommand.Arg(0), grepCommand.Arg(1))
		if err != nil {
			fmt.Println("cannot parse file '" + grepCommand.Arg(0) + "'")
			os.Exit(1)
		}

		result, err := a.Greper.Grep(rows, pattern)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		a.Printer.Print(result)
	}
}
