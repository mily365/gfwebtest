package service

import (
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	EnumCatalogSve *enumCatalogSve
)

type enumCatalogSve struct {
	gmeta.Meta `path:"service.enumCatalog"`
	base.ServiceBase
}

func init() {
	EnumCatalogSve = &enumCatalogSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(EnumCatalogSve)
}
