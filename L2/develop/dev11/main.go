package main

import (
	"github.com/mikhailbuslaev/wb-tasks/l2/dev11/server"
)

func main() {
	s := server.Server{}
	s.LogFileName = "server.log"
	s.Port = ":8090"
	s.ListenAndServe()
}