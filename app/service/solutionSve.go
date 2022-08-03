package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmeta"
	"github.com/gogf/gf/util/gutil"
	"os"
	"xpass/app"
	"xpass/app/service/base"
)

type FieldType string

const (
	Id       FieldType = "id"
	Int      FieldType = "int"
	Varchar  FieldType = "varchar"
	DateTime FieldType = "datetime"
)

func buildCreateSql(inputMap map[string]interface{}) string {
	for k, v := range inputMap {
		g.Dump(k, "==", v)
		if k == "validatorType" {
			f := gstr.Split(v.(string), "_")[0]
			if f == "required" {
				inputMap["NotNull"] = "NOT NULL"
			} else {
				inputMap["NotNull"] = "NULL"
			}
		}

	}
	sqlOneStr := makeSqlOneStrByFieldType(inputMap)
	return sqlOneStr
}

func makeSqlOneStrByFieldType(maps map[string]interface{}) string {
	rtn := ""
	fieldType := maps["sqlType"].(string)
	fieldType = gstr.Split(fieldType, "_")[0]
	switch fieldType {
	case "id":
		rtn = `${propName} int(${sqlLength}) unsigned ${NotNull} AUTO_INCREMENT`
	case "datetime":
		rtn = `${propName} timestamp ${NotNull} ${sqlDefault}`
	case "int":
		rtn = `${propName} ${sqlType}(${sqlLength}) ${NotNull}  ${sqlDefault}`
	case "tinyint":
		rtn = `${propName} ${sqlType}(1) ${NotNull}  ${sqlDefault}`
	case "varchar":
		rtn = `${propName} ${sqlType}(${sqlLength}) ${NotNull} ${sqlDefault}`
	case "json":
		rtn = `${propName} ${sqlType} ${NotNull} ${sqlDefault}`
	}
	singleStr := os.Expand(rtn, func(s string) string {
		g.Dump(s, maps)
		rtnStr := gconv.String(maps[s])
		if gstr.Contains(rtnStr, "_") {
			rtnStr = gstr.Split(rtnStr, "_")[0]
		}
		if s == "sqlDefault" {
			if rtnStr != "" {
				rtnStr = fmt.Sprintf("DEFAULT %s", rtnStr)
			} else {
				if fieldType == "varchar" && maps["NotNull"] != "NULL" {
					rtnStr = "DEFAULT '' "
				} else {
					rtnStr = "DEFAULT NULL"
				}

			}
		}
		return rtnStr
	})
	return singleStr
}

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
	errors, rtn := app.ModelFactory.TxModelActions(searchTable, func(tx *gdb.TX, model *gdb.Model) (error, interface{}) {
		//利用方案的bizCode作为表名
		//约定按照appName_sulutionName
		sl, _ := model.FindOne(i)
		g.Dump(sl.Map()["biz_code"])
		searchTable2 := g.Config().Get(app.ModelToTbl + "." + "ControlInfo").(string)
		cm := app.ModelFactory.GetModel(searchTable2)
		sp := app.TypePointerFuncFactory.GetStructArrayPointer("controlInfo")
		_ = cm.Where(map[string]interface{}{"sid": sl.Map()["id"]}).Scan(sp)
		//buildCreateSql(sp)
		bys, _ := gjson.Encode(sp)
		//mps := gconv.Map(bys)
		//m := gconv.Map(sp)
		mps := gconv.SliceMap(bys)
		var rtnTmps []string
		rtnTmps = append(rtnTmps, fmt.Sprintf("CREATE TABLE %s ( id int(11) NOT NULL AUTO_INCREMENT", sl.Map()["biz_code"].(string)))
		for _, v := range mps {
			sqlStr := buildCreateSql(v)
			if !gutil.IsEmpty(sqlStr) {
				rtnTmps = append(rtnTmps, sqlStr)
			}
		}
		rtnTmps = append(rtnTmps, "created_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP")
		rtnTmps = append(rtnTmps, "updated_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP")
		rtnTmps = append(rtnTmps, "deleted_time timestamp  NULL DEFAULT NULL")
		rtnTmps = append(rtnTmps, "PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
		sqlCreateStr := gstr.Join(rtnTmps, ",\n")
		g.Dump(sqlCreateStr, "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
		res, er := g.DB().Exec(sqlCreateStr)
		//生成表以后在实体文件内生成实体代码
		return er, res
	})
	if errors != nil {
		panic(errors)
	}
	return rtn
}

//再添加ID
//inputMap["SqlType"] = "id"
//inputMap["PropName"] = "id"
//inputMap["SqlLength"] = "10"
//inputMap["NotNull"] = "NOT NULL"
//inputMap["SqlDefault"] = ""
//sqlOneStr := makeSqlOneStrByFieldType(inputMap)
//rtn = append(rtn, sqlOneStr)
