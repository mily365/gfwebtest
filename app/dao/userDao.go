package dao

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/dao/base"
)

var (
	UserDao *userDao
)

type userDao struct {
	gmeta.Meta `path:"dao.user"`
	base.DaoBase
}

func init() {
	UserDao = &userDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(UserDao)
}
