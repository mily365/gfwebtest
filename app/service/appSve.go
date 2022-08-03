package service

import (
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	AppSve *appSve
)

type appSve struct {
	gmeta.Meta `path:"service.app"`
	base.ServiceBase
}

func init() {
	AppSve = &appSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(AppSve)
}
