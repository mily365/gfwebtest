package base

import (
	"gfwebtest/app"
	"github.com/gogf/gf/util/gmeta"
)

type adapterApi struct {
	gmeta.Meta `path:"api.*"`
	ApiBase
}

var (
	AdapterApi *adapterApi
)

func init()  {
	AdapterApi =&adapterApi{gmeta.Meta{}, ApiBase{}}
	app.AppContext.RegisterObj(AdapterApi)
}







