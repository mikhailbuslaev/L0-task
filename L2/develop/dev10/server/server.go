package server

import (
	"github.com/reiver/go-telnet"
	"fmt"
)

type Server struct {
	telnet.Server
}

func (s *Server) Set(address string) {
	s.Addr = address
	s.Handler = telnet.EchoHandler
}

func (s *Server) Run() error{
	fmt.Println("server run...")
	err := s.ListenAndServe()
	if nil != err {
		return err
	}
	return nil
}