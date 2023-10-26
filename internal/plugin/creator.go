package plugin

import (
	"gitea.com/ruijc/app/internal/core"
	"github.com/pangum/grpc"
)

type Creator struct {
	// 解决命名空间问题
}

func (c *Creator) Connection(client *grpc.Client) *core.Connection {
	return client.Connection("app")
}

func (c *Creator) Application(connection *core.Connection) *core.Application {
	return core.NewApplication(connection)
}

func (c *Creator) Config(connection *core.Connection) *core.Config {
	return core.NewConfig(connection)
}

func (c *Creator) Notification(connection *core.Connection) *core.Notification {
	return core.NewNotification(connection)
}
