package api

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
)

type enumCatalogApi struct {
	gmeta.Meta `path:"api.enumCatalog"`
	base.ApiBase
}

var (
	EnumCatalogApi *enumCatalogApi
)

func init() {
	EnumCatalogApi = &enumCatalogApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(EnumCatalogApi)
}
