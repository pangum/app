package app

import (
	"github.com/pangum/app/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		creator.Connection,
		creator.Application,
		creator.Notification,
		creator.Config,
	).Build().Build().Apply()
}
