package api

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
)

type userApi struct {
	gmeta.Meta `path:"api.user"`
	base.ApiBase
}

var (
	UserApi *userApi
)

func init() {
	UserApi = &userApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(UserApi)
}
