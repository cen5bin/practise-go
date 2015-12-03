package main

import (
	"fmt"
	"practise-go/client"
	"practise-go/server"
	//	"sync"
	"testing"
	"time"
)

func testTime(t *testing.T) {
	start := time.Now().UnixNano()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	end := time.Now().UnixNano()
	fmt.Println((end-start)/1000000, "ms")
}

func TestNet(t *testing.T) {
	port := 12222
	go func() {
		ser := &server.Server{}
		ser.Listen(port)
	}()
	time.Sleep(3 * time.Second)
	c := &client.Client{}
	c.Connect(fmt.Sprintf("127.0.0.1:%d", port))
	fmt.Println("start to test")
	start := time.Now().UnixNano() / 1000000
	cnt := 100000
	buf := []byte("aaaaaaassssssssssssssssssssssssssssssssssssssssss")
	for i := 0; i < cnt; i++ {
		c.Write(buf)
	}

	end := time.Now().UnixNano() / 1000000
	fmt.Printf("send %d times, cost %d ms\n", cnt, end-start)

}
