package client

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func (c *Client) Connect(addr string) {
	vals := strings.Split(addr, ":")
	if len(vals) != 2 {
		fmt.Println("err: addr must be like 127.0.0.1:8080")
		return
	}
	var err error
	c.conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("connect err, ", err)
		return
	}
}

func (c *Client) Write(buf []byte) {
	_, err := c.conn.Write(buf)
	if err != nil {
		fmt.Println("write err, ", err)
		return
	}
}

func (c *Client) Read(buf []byte) (n int) {
	var err error
	n, err = c.conn.Read(buf)
	if err != nil {
		fmt.Println("read error,", err)
	}
	return
}

func (c *Client) Close() {
	c.conn.Close()
}
