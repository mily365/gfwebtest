package base

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
)

var (
	AdapterDao *adapterDao
)

type adapterDao struct {
	gmeta.Meta `path:"dao.*"`
	DaoBase
}

func init() {
	AdapterDao = &adapterDao{gmeta.Meta{}, DaoBase{}}
	app.AppContext.RegisterObj(AdapterDao)
}
