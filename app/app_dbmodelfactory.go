package app

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
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
	m := mf.GetOrSet(tblName, g.DB().Model(tblName)).(*gdb.Model)
	err, rt := fun(tx, m.Clone())
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
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
	fp := TypePointerFuncFactory.RegisterOrGetTypePointer(explicitName, nil)
	up := fp()
	return up
}
func (tpf *typePointerFuncFactory) GetStructArrayPointer(explicitName string) interface{} {
	fp := TypePointerFuncFactory.RegisterOrGetTypePointer(explicitName+"s", nil)
	up := fp()
	return up
}

func (tpf *typePointerFuncFactory) GetFuncMapForModelPointer() map[string]func() interface{} {
	return tpf.mapFuncPointer
}
