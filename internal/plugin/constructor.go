package plugin

import (
	"github.com/pangum/app/internal/core"
	"github.com/pangum/app/internal/internal/constant"
	"github.com/pangum/grpc"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) Connection(client *grpc.Client) *core.Connection {
	return client.Connection(constant.LabelApp)
}

func (*Constructor) Application(connection *core.Connection) *core.Application {
	return core.NewApplication(connection)
}

func (*Constructor) Config(connection *core.Connection) *core.Config {
	return core.NewConfig(connection)
}

func (*Constructor) Notification(connection *core.Connection) *core.Notification {
	return core.NewNotification(connection)
}
