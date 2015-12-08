package server

import (
	"practise-go/net/client"
)

type EventHandler interface {
	OnClientConnected(client *client.Client)
}
