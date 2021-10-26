
package dao

import (
	"gfwebtest/app"
	"gfwebtest/app/dao/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	UserDetailDao *userDetailDao
)
type userDetailDao struct {
	gmeta.Meta `path:"dao.userDetail"`
	base.DaoBase
}
func init() {
	UserDetailDao=&userDetailDao{gmeta.Meta{},base.DaoBase{}}
	app.AppContext.RegisterObj(UserDetailDao)
}
