package base

import (
	"gfwebtest/app"
	"github.com/gogf/gf/util/gmeta"
)

var(
	AdapterSve *adapterSve

)
func init() {
	AdapterSve =&adapterSve{gmeta.Meta{}, ServiceBase{}}
	app.AppContext.RegisterObj(AdapterSve)

}
type adapterSve struct {
	gmeta.Meta `path:"service.*"`
	ServiceBase
}

