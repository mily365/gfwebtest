package base

import (
	"gfwebtest/app"
	"github.com/gogf/gf/util/gmeta"
)
var(
	AdapterDao *adapterDao
)
type adapterDao struct {
	gmeta.Meta `path:"dao.*"`
	DaoBase
}

func init() {
	AdapterDao =&adapterDao{gmeta.Meta{}, DaoBase{}}
	app.AppContext.RegisterObj(AdapterDao)
}