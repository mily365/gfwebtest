package service

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	SolutionSve *solutionSve
)

type solutionSve struct {
	gmeta.Meta `path:"service.solution"`
	base.ServiceBase
}

func init() {
	SolutionSve = &solutionSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(SolutionSve)
}

func (s *solutionSve) CreateTable(ctx context.Context, i interface{}) interface{} {
	g.Dump(".................", i)
	//按照id查询出元数据
	modelName := app.GetModelName(ctx, nil)
	//modelKey := gstr.CaseCamelLower(modelName)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	_, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		//利用方案的bizCode作为表名
		//约定按照appName_sulutionName
		sl, _ := model.FindOne(i)
		searchTable2 := g.Config().Get(app.ModelToTbl + "." + "ControlInfo").(string)
		cm := app.ModelFactory.GetModel(searchTable2)
		sp := app.TypePointerFuncFactory.GetStructArrayPointer("controlInfo")
		_ = cm.Where(map[string]interface{}{"sid": sl.Map()["id"]}).Scan(sp)
		g.Dump(sp)
		return nil, nil
	})

	return rtn
}
