package app

import (
	"context"

	"gitea/ruijc/storage/internal"

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

func (c *Config) Sms(ctx context.Context, appid int64) (string, error) {
	return c.get(ctx, appid, core.ConfigType_CONFIG_TYPE_SMS)
}

func (c *Config) Bucket(ctx context.Context, appid int64) (string, error) {
	return c.get(ctx, appid, core.ConfigType_CONFIG_TYPE_BUCKET)
}

func (c *Config) get(ctx context.Context, appid int64, typ core.ConfigType) (value string, err error) {
	req := new(config.GetByAppReq)
	req.App = appid
	req.Type = typ
	if rsp, ce := c.client.GetByApp(ctx, req); nil != ce {
		err = ce
	} else {
		value = rsp.Config.Value
	}

	return
}
