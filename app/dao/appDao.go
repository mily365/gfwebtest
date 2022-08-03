package dao

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/dao/base"
)

var (
	AppDao *appDao
)

type appDao struct {
	gmeta.Meta `path:"dao.app"`
	base.DaoBase
}

func init() {
	AppDao = &appDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(AppDao)
}
