package client

import (
	"github.com/reiver/go-telnet"
	"fmt"
)

type Client struct {
	Conn *telnet.Conn
}

func (c *Client) Connect(address string) error {	
	fmt.Println("client connect...")
	conn, err := telnet.DialTo(address)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

func (c *Client) Listen() error{
	bigBuf := make([]byte, 0, 100)
	buf := make([]byte, 1)
	for {
		_, err := c.Conn.Read(buf)
		if err != nil {
			return err
		}
		if buf[0] == 10 {
			fmt.Println("message received: "+string(bigBuf))
			bigBuf = make([]byte, 1)
		} else {
			bigBuf = append(bigBuf, buf...)
		}
	}
}

func (c *Client) Write(message []byte) {
	c.Conn.Write([]byte(message))
}