package app

import (
	"context"

	"gitea.com/ruijc/app/internal"

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

func (a *Application) Name(id Id) (name string, err error) {
	req := new(application.GetReq)
	req.Id = int64(id)
	if rsp, ce := a.client.Get(context.Background(), req); nil != ce {
		err = ce
	} else {
		name = gox.If(nil != rsp.Application, rsp.Application.Name)
	}

	return
}
