package api

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
)

type controlInfoApi struct {
	gmeta.Meta `path:"api.controlInfo"`
	base.ApiBase
}

var (
	ControlInfoApi *controlInfoApi
)

func init() {
	ControlInfoApi = &controlInfoApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(ControlInfoApi)
}
