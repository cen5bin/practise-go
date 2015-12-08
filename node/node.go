package node

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type Node struct {
	name     string
	port     int
	conn     net.Conn
	listener net.Listener
}

func (this *Node) listen(port int) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("listen error: ", err)
		return err
	}
	this.port = port
	this.listener = listener
	fmt.Println("listen on port", this.port)
	return nil
}

func (this *Node) Listen() error {
	for {
		port := int(time.Now().UnixNano()%10000) + 10000
		if this.listen(port) != nil {
			continue
		}
		break
	}
	return nil
}

func (this *Node) ListenOnPort(port int) error {
	return this.listen(port)
}

func (this *Node) onAccept(c net.Conn) {
	tcpConn := c.(*net.TCPConn)
	fmt.Println(tcpConn.RemoteAddr())
}

func (this *Node) Start() {
	go func() {
		for {
			c, err := this.listener.Accept()
			if err != nil {
				fmt.Println("accept error:", err)
				break
			}
			this.onAccept(c)
		}
	}()
}
