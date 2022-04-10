package dao

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/dao/base"
)

var (
	SolutionDao *solutionDao
)

type solutionDao struct {
	gmeta.Meta `path:"dao.solution"`
	base.DaoBase
}

func init() {
	SolutionDao = &solutionDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(SolutionDao)
}
