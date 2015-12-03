package main

import (
	"fmt"
	"practise-go/server"
)

func main() {
	server := &server.Server{}
	server.Listen(12222)
	fmt.Println(server)
	fmt.Println("asd")
}
