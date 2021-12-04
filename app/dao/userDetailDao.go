package dao

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/dao/base"
)

var (
	UserDetailDao *userDetailDao
)

type userDetailDao struct {
	gmeta.Meta `path:"dao.userDetail"`
	base.DaoBase
}

func init() {
	UserDetailDao = &userDetailDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(UserDetailDao)
}
