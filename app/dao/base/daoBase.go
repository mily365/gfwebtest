package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmeta"
	"github.com/gogf/gf/util/gutil"
	"reflect"
	"sort"
	"xpass/app"
)

type DaoBase struct {
}

//返回CaseCamel的实体名称
func getModelName(ctx context.Context, search g.Map) string {
	var modelName string = ""
	//路径映射优先
	path2ModelRegKey := ctx.Value(app.Path2ModelRegKey).(string)
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
func getModelAndTableName(ctx context.Context, search g.Map) (string, string) {
	var modelName string = ""
	//路径映射优先
	path2ModelRegKey := ctx.Value(app.Path2ModelRegKey).(string)
	if search != nil && search["model"] != nil {
		modelName = search["model"].(string)
	} else {
		modelName = gstr.CaseCamel(path2ModelRegKey)
	}
	if gutil.IsEmpty(modelName) {
		panic("please input model query param..")
	}
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName)
	if searchTable == nil {
		panic("please config model2table..")
	}
	return modelName, searchTable.(string)
}
func (s *DaoBase) Scrollpage(ctx context.Context, i interface{}) interface{} {
	rtn := new(app.SearchResult)
	search := i.(g.Map)
	//main model
	modelName := getModelName(ctx, search)
	modelKey := gstr.CaseCamelLower(modelName)
	app.Logger.Debug("dao scrollpage....", modelKey)
	res := app.GetEsFactory().ScrollPage(ctx, search, modelKey)

	rtn.ScrollId = res.ScrollId

	//var rows []interface{}
	//for _, hit := range res.Hits.Hits {
	//	sp := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	//	err := json.Unmarshal(hit.Source, &sp)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	rows = append(rows, sp)
	//}
	rtn.Rows = app.TranEsResultToRows(res, modelKey)
	return rtn
}
func (s *DaoBase) Withalls(ctx context.Context, i interface{}) interface{} {
	rtn := new(app.SearchResult)
	search := i.(g.Map)
	//main model
	modelName := getModelName(ctx, search)
	modelKey := gstr.CaseCamelLower(modelName)

	//

	//fetch combined entity
	sp := app.TypePointerFuncFactory.GetStructArrayPointer(modelKey)
	metadata := gmeta.Data(sp)
	if metadata == nil {
		panic("defind entity with relation metaData!")
	}
	g.Dump(metadata)
	keys := make([]string, 0)
	for role := range metadata {
		keys = append(keys, role)
	}
	sort.Strings(keys)
	//gmeta.Meta `master:"User_uid#" second:"UserDetail_uid#User_Uid" third:"UserScore_uid#UserDetail_Uid"`
	//gmeta.Meta `a:"User_id#" b:"UserDetail_uid#User_Id" c:"UserScore_uid#User_Id"`
	for _, role := range keys {
		g.Log().Debug(role + "ccccccccccccccccccccccc")
		rel := metadata[role]
		relArray := gstr.Split(rel.(string), "#")
		if len(relArray) != 2 {
			panic("please define entity with relations metadata!")
		}
		modelnameTablekey := relArray[0]
		modelnameTablekeyArray := gstr.Split(modelnameTablekey, "_")
		if len(modelnameTablekeyArray) != 2 {
			panic("check entity with relation metadata modelnameTablekey!")
		}
		modelname := modelnameTablekeyArray[0]
		tablekey := modelnameTablekeyArray[1]

		var orderByStrings = []string{""}
		if search["orderBy"] != nil {
			orderBy := search["orderBy"].([]interface{})
			//orderByStrings = gconv.SliceStr(orderBy)
			orderByStrings = s.buildOrderBy(gconv.SliceStr(orderBy), gstr.CaseCamelLower(modelname))
		}

		searchTable := g.Config().Get("model2Tbl." + modelname)

		if gutil.IsEmpty(searchTable) {
			panic("please config model2tabl for " + modelname + " in config file!")
		}
		if role == "a" {
			skip, pageSize := app.PageParam(search)
			um := app.ModelFactory.GetModel(searchTable.(string))
			var cntM = um.Clone()
			g.Dump(gstr.CaseCamelLower(modelname), "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			//whereForm := s.buildWhereSqlMapByInputMap(search["queryForm"].(g.Map), gstr.CaseCamelLower(modelname))
			if search["fields"] != nil && gstr.Trim(search["fields"].(string)) != "" {
				err := s.buildWhereLikeModelByInputMap(search["queryForm"].(g.Map), gstr.CaseCamelLower(modelname), um.Fields(search["fields"])).Offset(skip).Limit(pageSize).Order(orderByStrings...).ScanList(sp, modelname)
				//err := um.Fields(search["fields"]).Where(whereForm).Offset(skip).Limit(pageSize).Order(orderByStrings...).ScanList(sp, modelname)
				if err != nil {
					panic(err.Error())
				}
			} else {
				err := s.buildWhereLikeModelByInputMap(search["queryForm"].(g.Map), gstr.CaseCamelLower(modelname), um.Fields()).Offset(skip).Limit(pageSize).Order(orderByStrings...).ScanList(sp, modelname)
				if err != nil {
					panic(err.Error())
				}
			}

			//cnt, e := cntM.Count(whereForm)
			cnt, e := s.buildWhereLikeModelByInputMap(search["queryForm"].(g.Map), gstr.CaseCamelLower(modelname), cntM).Count()
			if e != nil {
				panic(e.Error())
			}
			rtn.Total = cnt
		} else {
			modelrel := relArray[1]
			//gmeta.Meta `a:"User_id#" b:"UserDetail_uid#User_Id" c:"UserScore_uid#User_Id"`
			depEntity := gstr.Split(modelrel, "_")[0]
			depEntityProperty := gstr.Split(modelrel, "_")[1]
			//confirm special table
			um2 := app.ModelFactory.GetModel(searchTable.(string)).Clone()
			g.Log().Debug(depEntity + "---" + depEntityProperty + "tablekey---" + tablekey)
			if search["subFields"] != nil {
				um2.Fields(search["subFields"].(g.Map)[modelname]).Where(tablekey, gdb.ListItemValuesUnique(sp, depEntity, depEntityProperty)).
					ScanList(sp, modelname, depEntity, tablekey+":"+depEntityProperty)
			} else {
				um2.Fields().Where(tablekey, gdb.ListItemValuesUnique(sp, depEntity, depEntityProperty)).
					ScanList(sp, modelname, depEntity, tablekey+":"+depEntityProperty)
			}

		}
	}
	rtn.Rows = sp
	return rtn
}
func (s *DaoBase) buildWhereSqlMapByInputMap(search g.Map, modelKey string) g.Map {
	rtn := g.Map{}
	searchForm := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	structType := reflect.TypeOf(searchForm).Elem().Elem()
	for k, v := range search {
		if v != nil {
			//首字母大写
			upperFirstCh := gstr.CaseCamel(k)
			sf, isFind := structType.FieldByName(upperFirstCh)
			if isFind == true {
				tmpValue := sf.Tag.Get("orm")
				rtn[tmpValue] = fmt.Sprintf("%s%s%s", "%", v, "%")
			}
		}
	}
	return rtn
}

//
func (s *DaoBase) buildOrderBy(orderByStr []string, modelKey string) []string {
	var rtn []string
	searchForm := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	structType := reflect.TypeOf(searchForm).Elem().Elem()
	for _, orderStr := range orderByStr {
		g.Dump(orderStr)
		inputProp := gstr.Split(orderStr, " ")[0]
		g.Dump(inputProp)
		upperFirstCh := gstr.CaseCamel(inputProp)
		sf, isFind := structType.FieldByName(upperFirstCh)
		if isFind == true {
			tmpValue := sf.Tag.Get("orm")
			g.Dump(inputProp)
			if inputProp != "id" {
				orderStr = gstr.Replace(orderStr, inputProp, tmpValue)
			}

			rtn = append(rtn, orderStr)
		}
	}
	g.Dump(rtn)
	return rtn
}

// 针对传入的查询条件进行orm查询条件的转换
func (s *DaoBase) buildWhereLikeModelByInputMap(search g.Map, modelKey string, mdl *gdb.Model) *gdb.Model {
	searchForm := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	structType := reflect.TypeOf(searchForm).Elem().Elem()
	for k, v := range search {
		if v != nil {
			//首字母大写
			upperFirstCh := gstr.CaseCamel(k)
			sf, isFind := structType.FieldByName(upperFirstCh)
			if isFind == true {
				tmpValue := sf.Tag.Get("orm")
				if tmpValue == "id" {
					mdl.Where(tmpValue, v)
				} else {
					mdl.WhereLike(tmpValue, fmt.Sprintf("%s%s%s", "%", v, "%"))
				}
			}
		}

	}
	return mdl
}
func (s *DaoBase) All(ctx context.Context, i interface{}) interface{} {
	app.Logger.Debug("dao all called......")
	rtn := new(app.SearchResult)
	search := i.(g.Map)
	modelName, searchTable := getModelAndTableName(ctx, search)
	//全部变成小写--计划
	modelKey := gstr.CaseCamelLower(modelName)
	if g.Config().GetBool("appInfo.enableEs") == true {
		esRes := app.GetEsFactory().All(ctx, search, modelKey)
		rtn.Total = int(esRes.Hits.TotalHits.Value)
		rtn.Rows = app.TranEsResultToRows(esRes, modelKey)
		return rtn
	}

	var sp interface{}
	var um *gdb.Model
	//约定类型函数工厂的key取实体的首字母小写
	sp = app.TypePointerFuncFactory.GetStructArrayPointer(modelKey)
	um = app.ModelFactory.GetModel(searchTable)
	var cntM = um.Clone()
	skip, pageSize := app.PageParam(search)
	var orderByStrings = []string{""}
	if search["orderBy"] != nil {
		orderBy := search["orderBy"].([]interface{})
		orderByStrings = s.buildOrderBy(gconv.SliceStr(orderBy), modelKey)
	}
	//whereForm := s.buildWhereSqlMapByInputMap(search["queryForm"].(g.Map), modelKey)
	err := s.buildWhereLikeModelByInputMap(search["queryForm"].(g.Map), modelKey, um.Fields(search["fields"])).OmitEmptyWhere().WhereNull("deleted_time").Offset(skip).Limit(pageSize).Order(orderByStrings...).Scan(sp)
	cnt, err := s.buildWhereLikeModelByInputMap(search["queryForm"].(g.Map), modelKey, cntM).WhereNull("deleted_time").Count()
	//cnt, err := cntM.OmitEmptyWhere().Count()
	if sp == nil {
		panic(err.Error())
	}

	rtn.Rows = sp
	rtn.Total = cnt
	return rtn
}

func (s *DaoBase) Create(ctx context.Context, i interface{}) interface{} {
	g.Log().Debug("create..............................................................")
	modelName := getModelName(ctx, nil)
	modelKey := gstr.CaseCamelLower(modelName)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)

	um := app.ModelFactory.GetModel(searchTable)
	um2 := um.Clone()
	rid, err := um.Data(i).InsertAndGetId()
	if err != nil {
		panic(err.Error())
	}
	rt, e := um2.Where(g.Map{"id": rid}).One()

	if e != nil {
		panic(e.Error())
	}
	//mp := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	if g.Config().GetBool("appInfo.enableEs") == true {
		app.GetEsFactory().Create(ctx, gconv.String(rid), rt.Json(), modelKey)
	}

	return rt
}
func (s *DaoBase) Createtx(ctx context.Context, i interface{}, tx *gdb.TX, model *gdb.Model) interface{} {
	g.Log().Debug("dao create..............................................................")
	um2 := model.Clone()
	rid, err := model.TX(tx).Data(i).InsertAndGetId()
	if err != nil {
		panic(err.Error())
	}
	rt, e := um2.TX(tx).Where(g.Map{"id": rid}).One()
	if e != nil {
		panic(e.Error())
	}
	//mp := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	return rt
}

func (s *DaoBase) CreateEsTx(ctx context.Context, rt interface{}, modelKey string) {
	if g.Config().GetBool("appInfo.enableEs") == true {
		app.GetEsFactory().Create(ctx, gconv.String(rt.(gdb.Record).Map()["id"]), rt.(gdb.Record).Json(), modelKey)
	}

}

func (s *DaoBase) Update(ctx context.Context, i interface{}) interface{} {
	modelName := getModelName(ctx, nil)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)
	_, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		um2 := model.Clone()
		if i.(g.Map)["id"] == nil {
			panic("lack id for update!")
		}
		if i.(g.Map)["version"] == nil {
			panic("lack version for update!")
		}
		idvalue := i.(g.Map)["id"]
		incversionPrev := gconv.Int(i.(g.Map)["version"].(json.Number).String())
		incversion := incversionPrev + 1
		i.(g.Map)["version"] = incversion
		delete(i.(g.Map), "id")
		rt, _ := um2.TX(tx).Where(g.Map{"id": idvalue, "version>=": incversion}).LockShared().One()
		if rt != nil {
			g.Log().Debug("update concurrent happened...!")
			return nil, rt
		}
		_, err := model.TX(tx).Update(i.(g.Map), g.Map{"id": idvalue, "version<": incversion})
		if err != nil {
			panic(err.Error())
		}
		return nil, nil
	})
	g.Dump(rtn)
	return rtn
}

func (s *DaoBase) Updatetx(ctx context.Context, i interface{}, tx *gdb.TX, model *gdb.Model) interface{} {
	um2 := model.Clone()
	if i.(g.Map)["id"] == nil {
		panic("lack id for update!")
	}
	if i.(g.Map)["version"] == nil {
		panic("lack version for update!")
	}
	idvalue := i.(g.Map)["id"]
	incversionPrev := gconv.Int(i.(g.Map)["version"].(json.Number).String())
	incversion := incversionPrev + 1
	i.(g.Map)["version"] = incversion
	delete(i.(g.Map), "id")
	rt, _ := um2.TX(tx).Where(g.Map{"id": idvalue, "version>=": incversion}).LockShared().One()
	if rt != nil {
		g.Log().Debug("update concurrent happened...!")
		return nil
	}
	rtn2, err := model.TX(tx).Update(i.(g.Map), g.Map{"id": idvalue, "version<": incversion})
	if err != nil {
		panic(err.Error())
	}
	return rtn2
}

func (s *DaoBase) Delete(ctx context.Context, i interface{}) interface{} {
	modelName := getModelName(ctx, nil)
	searchTable := g.Config().Get(app.ModelToTbl + "." + modelName).(string)

	um := app.ModelFactory.GetModel(searchTable)
	//rtn, err := um.Where("id in (?)", i.(g.Map)["ids"]).Delete()
	m := map[string]interface{}{"deleted_time": gtime.Now()}
	rtn, err := um.Where("id in (?)", i.(g.Map)["ids"]).Update(m)
	if err != nil {
		panic(err.Error())
	}
	g.Dump(rtn)
	return rtn
}
func (s *DaoBase) Deletetx(ctx context.Context, i interface{}, tx *gdb.TX, model *gdb.Model) interface{} {
	return nil
}
