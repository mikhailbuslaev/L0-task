package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
	ps "github.com/mitchellh/go-ps"
	"time"
)

var (
	noPathError error = errors.New("path required")
	noArgumentError error = errors.New("argument required")
)

func handleShell(input string, done chan struct{}) error {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.TrimSuffix(input, "\n")
	inputArr := strings.Split(input, "|")
	for i := range inputArr {
		args := strings.Split(inputArr[i], " ")
		switch args[0] {
		case "quit":
			fmt.Println("programm goes sleep...")
			time.Sleep(1 * time.Second)
			done <- struct{}{}
		case "cd":
			if len(args) < 2 {
				return  noPathError
			}
			return os.Chdir(args[1])
		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		case "echo":
			if len(args) < 2 {
				return  noArgumentError
			}
			fmt.Println(args[1])
		case "kill":
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			process, err := os.FindProcess(pid);
			if err != nil {
				return err
			}
			err = process.Kill()
			if err != nil {
				return err
			}
		case "ps":
			processes, err := ps.Processes()
			if err != nil {
				return err
			}
			for _, v := range processes {
				fmt.Println(v.Pid(), v.Executable())
			}
		// case "fork":
		// 	path, err := os.Getwd()
		// 	if err != nil {
		// 		return err
		// 	}
		// 	newProcess, err := os.StartProcess("newprocess", args[1:], &os.ProcAttr{Dir:path})
		default:
			fmt.Println("hello from shell!")
		}
	}
	return nil
}

func main() {
	done := make(chan struct{})
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("dev08 shell started!")
	go func() {
		for {
			fmt.Print("> ")
			// Read the keyboad input.
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			// Handle the execution of the input.
			if err = handleShell(input, done); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}()
	<-done
	fmt.Println("exit shell ...")
}
