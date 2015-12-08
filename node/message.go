package node

import (
	"encoding/json"
	//	"fmt"
)

type Message map[string]interface{}

func NewMessage() Message {
	return make(map[string]interface{})
}

func ParseMessage(buf []byte) (Message, error) {
	var f interface{}
	err := json.Unmarshal(buf, &f)
	return f.(map[string]interface{}), err
}

func (m Message) Get(key string) interface{} {
	return m[key]
}

func (m Message) Set(key string, value interface{}) {
	m[key] = value
}

func (m Message) ToBytes() ([]byte, error) {
	return json.Marshal(m)
}
