package main

import (
	"github.com/mikhailbuslaev/wb-tasks/l2/dev10/server"
	"github.com/mikhailbuslaev/wb-tasks/l2/dev10/client"
	"time"
	"os"
	"io"
	"fmt"
	"flag"
)

var (
	addressFlag *string
	timeoutFlag *time.Duration //in seconds
)

func timer(timeout time.Duration) <- chan struct{}{
	timer := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		timer <- struct{}{}
	}()
	return timer
}

func runTLS(address string) {
	// server run
	server := &server.Server{}
	server.Set(address)
	go func() {
		if err := server.Run(); err != nil {
			err.Error()
			os.Exit(1)
		}
	}()
	// client run
	client := &client.Client{}
	if err := client.Connect(address); err != nil {
		err.Error()
		os.Exit(1)
	}
	// we need close conn before exit
	defer func() {
		fmt.Println("connenction close...")
		client.Conn.Close()
	}()
	// client listen
	go func() {
		if err := client.Listen(); err != nil {
			err.Error()
			os.Exit(1)
		}
	}()
	// timeout implementation
	timer := timer(*timeoutFlag)
	// write to server
	func() {
		for {
			select {
			case <- timer:
				return
			default:
				var message string
				fmt.Print("print here: ")
				_, err := fmt.Fscan(os.Stdin, &message)
				if err != nil {
					// ctrl+z shutdown
					if err == io.EOF {
						return
					}
					err.Error()
				}
				client.Write([]byte(message+"\n"))
				time.Sleep(1*time.Second)
			}
		}
	}()
	fmt.Println("programm exit...")
	time.Sleep(1*time.Second)
}

func main() {
	tlsCommand := flag.NewFlagSet("tls", flag.ExitOnError)
	addressFlag = tlsCommand.String("addr", ":1111", "reverse sorting order")
	timeoutFlag = tlsCommand.Duration("timeout", 10, "reverse sorting order")
	if len(os.Args) < 2 {
		fmt.Println("expected 'tls' subcommand")
		os.Exit(1)
	}

	if os.Args[1] == "tls" {
		tlsCommand.Parse(os.Args[2:])
		runTLS(tlsCommand.Arg(0))
	}
}
