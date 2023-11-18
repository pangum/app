package app

import (
	"github.com/pangum/app/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.Connection,
		ctor.Application,
		ctor.Notification,
		ctor.Config,
	).Build().Build().Apply()
}
