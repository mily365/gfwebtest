
package dao

import (
	"xpass/app"
	"xpass/app/dao/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	EnumCatalogDao *enumCatalogDao
)
type enumCatalogDao struct {
	gmeta.Meta `path:"dao.enumCatalog"`
	base.DaoBase
}
func init() {
	EnumCatalogDao=&enumCatalogDao{gmeta.Meta{},base.DaoBase{}}
	app.AppContext.RegisterObj(EnumCatalogDao)
}
