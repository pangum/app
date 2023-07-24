package app

import (
	"context"

	"gitea/ruijc/storage/internal"

	"github.com/goexl/gox"
	"gitlab.com/ruijc/app/application"
)

// Application 应用
type Application struct {
	client application.RpcClient
}

func newApplication(connection *internal.Connection) *Application {
	return &Application{
		client: application.NewRpcClient(connection),
	}
}

func (a *Application) Name(app int64) (name string, err error) {
	req := new(application.GetReq)
	req.Id = app
	if rsp, ce := a.client.Get(context.Background(), req); nil != ce {
		err = ce
	} else {
		name = gox.If(nil != rsp.Application, rsp.Application.Name)
	}

	return
}
