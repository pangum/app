package core

import (
	"context"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/structer"
	"gitlab.com/ruijc/app/core"
	"gitlab.com/ruijc/app/core/config"
	"gitlab.com/ruijc/app/core/kernel"
	"gitlab.com/ruijc/app/core/vo"
	"gitlab.com/ruijc/app/protocol"
)

// Config 配置
type Config struct {
	client protocol.ConfigClient
}

func NewConfig(connection *Connection) *Config {
	return &Config{
		client: protocol.NewConfigClient(connection),
	}
}

func (c *Config) Storage(ctx context.Context, id core.Id) (value *vo.Storage, err error) {
	value = new(vo.Storage)
	err = c.get(ctx, id, kernel.ConfigType_CONFIG_TYPE_STORAGE, value)

	return
}

func (c *Config) get(ctx context.Context, id core.Id, typ kernel.ConfigType, value any) (err error) {
	req := new(config.GetByAppReq)
	req.App = int64(id)
	req.Type = typ
	if rsp, ce := c.client.GetByApp(ctx, req); nil != ce {
		err = ce
	} else if nil != rsp.Config && nil != rsp.Config.Value {
		err = structer.Converter().From(rsp.Config.Value.AsMap()).To(value).Build().Convert()
	} else {
		err = exc.NewFields("找不到配置信息", field.New("id", id), field.New("type", typ.String()))
	}

	return
}
