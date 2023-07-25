package internal

import (
	"github.com/pangum/grpc"
)

type Connection = grpc.Connection

func newConnection(client *grpc.Client) *Connection {
	return client.Connection("app")
}
