package app

import (
	"gitea.com/ruijc/app/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.Connection,
		creator.Application,
		creator.Notification,
		creator.Config,
	).Build().Build().Apply()
}
