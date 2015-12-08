package node

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	n1 := &Node{}
	n1.Listen()
	n1.Start()
	fmt.Println("asd")
	time.Sleep(1000 * time.Second)
}
