package server

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	conn net.Conn
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			break
		}
		fmt.Printf("read %s", string(buf[:n]))
		conn.Write([]byte("asdsada"))
	}
}

func (s *Server) Listen(port int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err: ", err)
			break
		}
		go handleConn(c)
	}
}
