package app

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newApplication,
		newNotification,
		newConfig,
	)
}
