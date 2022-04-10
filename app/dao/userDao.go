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
	gmeta.Meta `path:"dao.user" table:"user"`
	base.DaoBase
}

//func (s *userDao) All(ctx context.Context, i interface{}) interface{} {
//	mt := gmeta.Data(s)
//	for m := range mt {
//		g.Dump(m)
//	}
//	return s.DaoBase.All(ctx, i)
//}

func init() {
	UserDao = &userDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(UserDao)
}
