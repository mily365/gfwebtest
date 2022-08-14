package dao

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/dao/base"
)

var (
	AppDao *appDao
)

type appDao struct {
	gmeta.Meta `path:"dao.app"`
	base.DaoBase
}

func init() {
	AppDao = &appDao{gmeta.Meta{}, base.DaoBase{}}
	app.AppContext.RegisterObj(AppDao)
}

func (application *appDao) FetchApp(ctx context.Context, i interface{}) interface{} {
	modelName := base.GetModelName(ctx, nil)
	modelKey := gstr.CaseCamelLower(modelName)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	um := app.ModelFactory.GetModel(searchTable)
	rtn, err := um.FindOne(application.BuildWhereSqlMapByInputMap(i.(g.Map), modelKey))
	if err != nil {
		panic(err)
	}
	return rtn
}
