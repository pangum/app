package app

import (
	"context"

	"gitea.com/ruijc/app/internal"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"

	"gitlab.com/ruijc/app/config"
	"gitlab.com/ruijc/app/core"
)

// Config 配置
type Config struct {
	client config.RpcClient
}

func newConfig(connection *internal.Connection) *Config {
	return &Config{
		client: config.NewRpcClient(connection),
	}
}

func (c *Config) Sms(ctx context.Context, id Id) (string, error) {
	return c.get(ctx, id, core.ConfigType_CONFIG_TYPE_SMS)
}

func (c *Config) Bucket(ctx context.Context, id Id) (string, error) {
	return c.get(ctx, id, core.ConfigType_CONFIG_TYPE_BUCKET)
}

func (c *Config) get(ctx context.Context, id Id, typ core.ConfigType) (value string, err error) {
	req := new(config.GetByAppReq)
	req.App = int64(id)
	req.Type = typ
	if rsp, ce := c.client.GetByApp(ctx, req); nil != ce {
		err = ce
	} else if nil != rsp.Config {
		value = rsp.Config.Value
	} else {
		err = exc.NewFields("找不到配置信息", field.New("id", id), field.New("type", typ.String()))
	}

	return
}
