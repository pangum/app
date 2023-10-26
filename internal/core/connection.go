package core

import (
	"github.com/pangum/grpc"
)

type Connection = grpc.Connection

func NewConnection(client *grpc.Client) *Connection {
	return client.Connection("app")
}
