
package dao

import (
	"gfwebtest/app"
	"gfwebtest/app/dao/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	UserDao *userDao
)
type userDao struct {
	gmeta.Meta `path:"dao.user"`
	base.DaoBase
}
func init() {
	UserDao=&userDao{gmeta.Meta{},base.DaoBase{}}
	app.AppContext.RegisterObj(UserDao)
}
