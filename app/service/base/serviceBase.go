package base

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"xpass/app"
)

type ServiceBase struct {
	Dao app.DaoOperation
}

func (s *ServiceBase) Withalls(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Withalls(ctx, i)
}
func (s *ServiceBase) Scrollpage(ctx context.Context, i interface{}) interface{} {
	g.Log().Debug("service Scrollpage")
	return s.Dao.Scrollpage(ctx, i)
}
func (s *ServiceBase) All(ctx context.Context, i interface{}) interface{} {
	g.Log().Debug("service ServiceBase  " + ctx.Value(app.Path2ModelRegKey).(string))
	return s.Dao.All(ctx, i)
}

func (s *ServiceBase) Create(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Create(ctx, i)
}
func (s *ServiceBase) Createtx(ctx context.Context, i interface{}) interface{} {
	modelName := app.GetModelName(ctx, nil)
	modelKey := gstr.CaseCamelLower(modelName)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	_, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		outRtn := s.Dao.Createtx(ctx, i, tx, model)
		s.Dao.CreateEsTx(ctx, outRtn, modelKey)
		return nil, outRtn
	})
	return rtn
}
func (s *ServiceBase) Update(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Update(ctx, i)
}

func (s *ServiceBase) Delete(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Delete(ctx, i)
}
func (s *ServiceBase) Updatetx(ctx context.Context, i interface{}) interface{} {
	modelName := app.GetModelName(ctx, nil)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	_, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		outRtn := s.Dao.Updatetx(ctx, i, tx, model)
		//s.Dao.CreateEsTx(ctx, outRtn, modelKey)
		return nil, outRtn
	})
	return rtn
}

func (s *ServiceBase) Deletetx(ctx context.Context, i interface{}) interface{} {
	modelName := app.GetModelName(ctx, nil)
	//modelKey := gstr.CaseCamelLower(modelName)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	_, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		outRtn := s.Dao.Deletetx(ctx, i, tx, model)
		//s.Dao.CreateEsTx(ctx, outRtn, modelKey)
		return nil, outRtn
	})
	return rtn
}
