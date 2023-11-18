package core

import (
	"context"

	"github.com/goexl/gox"
	"gitlab.com/ruijc/app/core"
	"gitlab.com/ruijc/app/core/application"
	"gitlab.com/ruijc/app/protocol"
)

// Application 应用
type Application struct {
	client protocol.ApplicationClient
}

func NewApplication(connection *Connection) *Application {
	return &Application{
		client: protocol.NewApplicationClient(connection),
	}
}

func (a *Application) Name(id core.Id) (name string, err error) {
	req := new(application.GetReq)
	req.Id = int64(id)
	if rsp, ce := a.client.Get(context.Background(), req); nil != ce {
		err = ce
	} else {
		name = gox.If(nil != rsp.Application, rsp.Application.Name)
	}

	return
}
