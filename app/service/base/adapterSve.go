package base

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
)

var (
	AdapterSve *adapterSve
)

func init() {
	AdapterSve = &adapterSve{gmeta.Meta{}, ServiceBase{}}
	app.AppContext.RegisterObj(AdapterSve)

}

type adapterSve struct {
	gmeta.Meta `path:"service.*"`
	ServiceBase
}
