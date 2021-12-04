package service

import (
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	UserSve *userSve
)

type userSve struct {
	gmeta.Meta `path:"service.user"`
	base.ServiceBase
}

func init() {
	UserSve = &userSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(UserSve)
}
