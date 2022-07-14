package app

import (
	"context"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gutil"
)

//---------------mysql的模型工厂-------------------------
type modelFactory struct {
	*gmap.StrAnyMap
}

func (mf *modelFactory) GetModel(tblName string) *gdb.Model {
	m := mf.GetOrSet(tblName, g.DB().Model(tblName)).(*gdb.Model)
	return m.Clone()
}
func (mf *modelFactory) TxModelActions(tblName string, fun func(tx *gdb.TX, model *gdb.Model) (error, interface{})) (error, interface{}) {
	tx, ex := g.DB().Begin()
	if ex != nil {
		panic(ex.Error())
	}
	//m := mf.GetOrSet(tblName, g.DB().Model(tblName)).(*gdb.Model)
	mClone := mf.GetModel(tblName)
	err, rt := fun(tx, mClone)
	defer func() {
		if err != nil {
			if er := tx.Rollback(); er != nil {
				panic(er.Error())
			}
		} else {
			if erx := tx.Commit(); erx != nil {
				panic(erx.Error())
			}
		}
	}()
	return err, rt
}

//------------------------
type typePointerFuncFactory struct {
	mapFuncPointer map[string]func() interface{}
}

func (tpf *typePointerFuncFactory) RegisterOrGetTypePointer(typeName string, typeFunPointer func() interface{}) func() interface{} {
	g.Log().Info(tpf.mapFuncPointer)
	if tpf.mapFuncPointer[typeName] != nil {
		return tpf.mapFuncPointer[typeName]
	} else {
		if typeFunPointer == nil {
			panic("please input param typeFunPointer...or check whether generate model register code!")
		}
		tpf.mapFuncPointer[typeName] = typeFunPointer
		return typeFunPointer
	}
}
func (tpf *typePointerFuncFactory) GetStructPointer(explicitName string) interface{} {
	fp := tpf.RegisterOrGetTypePointer(explicitName, nil)
	up := fp()
	return up
}
func (tpf *typePointerFuncFactory) GetStructArrayPointer(explicitName string) interface{} {
	fp := tpf.RegisterOrGetTypePointer(explicitName+"s", nil)
	up := fp()
	return up
}

func (tpf *typePointerFuncFactory) GetFuncMapForModelPointer() map[string]func() interface{} {
	return tpf.mapFuncPointer
}

//返回CaseCamel的实体名称
func GetModelName(ctx context.Context, search g.Map) string {
	var modelName string = ""
	//路径映射优先
	path2ModelRegKey := ctx.Value(Path2ModelRegKey).(string)
	if search != nil && search["model"] != nil {
		modelName = search["model"].(string)
	} else {
		modelName = gstr.CaseCamel(path2ModelRegKey)
	}
	if gutil.IsEmpty(modelName) {
		panic("please input model query param..")
	}

	return modelName
}

//返回CaseCamel的实体名称
func GetModelAndTableName(ctx context.Context, search g.Map) (string, string) {
	var modelName string = ""
	//路径映射优先
	path2ModelRegKey := ctx.Value(Path2ModelRegKey).(string)
	if search != nil && search["model"] != nil {
		modelName = search["model"].(string)
	} else {
		modelName = gstr.CaseCamel(path2ModelRegKey)
	}
	if gutil.IsEmpty(modelName) {
		panic("please input model query param..")
	}
	searchTable := g.Config().Get(ModelToTbl + "." + modelName)
	if searchTable == nil {
		panic("please config model2table..")
	}
	return modelName, searchTable.(string)
}
