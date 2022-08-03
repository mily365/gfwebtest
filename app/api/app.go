package api

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
)

type appApi struct {
	gmeta.Meta `path:"api.app"`
	base.ApiBase
}

var (
	AppApi *appApi
)

func init() {
	AppApi = &appApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(AppApi)
}
