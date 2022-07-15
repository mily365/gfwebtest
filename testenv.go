package main

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"os"
)

type x struct {
	M string `json:"m"`
	N string `json:"n"`
}

var tmpMap = map[string]interface{}{
	"text": "ddd",
}

type FieldType string

const (
	Id       FieldType = "id"
	Int      FieldType = "int"
	Varchar  FieldType = "varchar"
	DateTime FieldType = "datetime"
)

func buildCreateSql(inputMap map[string]interface{}) string {
	var rtn []string
	var mapTmp = make(map[string]string)
	for k, v := range inputMap {
		if k == "sqlType" {
			mapTmp["SqlType"] = v.(string)
		}
		if k == "propName" {
			mapTmp["PropName"] = v.(string)
		}
		if k == "sqlLength" {
			mapTmp["SqlLength"] = v.(string)
		}
		if k == "sqlDefault" {
			if v == "" {
				mapTmp["SqlDefault"] = "DEFAULT NULL"
			} else {
				mapTmp["SqlDefault"] = v.(string)
			}
		}
		if k == "validateType" {
			if v == "required" {
				mapTmp["NotNull"] = "NOT NULL"
			} else {
				mapTmp["NotNull"] = ""
			}
		}

		sqlOneStr := makeSqlOneStrByFieldType(mapTmp)
		rtn = append(rtn, sqlOneStr)

	}
	//再添加ID
	mapTmp["SqlType"] = "id"
	mapTmp["PropName"] = "id"
	mapTmp["SqlLength"] = "10"
	mapTmp["NotNull"] = "NOT NULL"
	mapTmp["SqlDefault"] = ""
	sqlOneStr := makeSqlOneStrByFieldType(mapTmp)
	rtn = append(rtn, sqlOneStr)
	fmt.Println(rtn)
	return gstr.Join(rtn, ",")
}

func makeSqlOneStrByFieldType(maps map[string]string) string {
	rtn := ""
	fieldType := FieldType(maps["SqlType"])
	switch fieldType {
	case Id:
		rtn = `${PropName} int(${SqlLength}) unsigned ${NotNull} AUTO_INCREMENT`
	case DateTime:
		rtn = `${PropName} ${SqlType} ${NotNull} ${SqlDefault}`
	case Int:
		rtn = `${PropName} ${SqlType}(${SqlLength}) ${NotNull}  ${SqlDefault}`
	case Varchar:
		rtn = `${PropName} ${SqlType}(${SqlLength}) ${NotNull} ${SqlDefault}`
	}
	singleStr := os.Expand(rtn, func(s string) string {
		return maps[s]
	})
	return singleStr
}
func main() {
	//	sqlTmpl := `
	//
	//CREATE TABLE ${TableName} (
	//
	//  name varchar(45) NOT NULL,
	//  created_at datetime DEFAULT NULL,
	//  updated_at datetime DEFAULT NULL,
	//  deleted_at datetime DEFAULT NULL,
	//  version int(11) DEFAULT 0,
	//  age int(11) DEFAULT NULL,
	//  PRIMARY KEY (${PrimaryKey})
	//) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
	//
	//
	//`
	//fmt.Println(genv.Get("GF_GCFG_PATH"))
	//m := g.MapStrAny{"hello": "ok"}
	//v := &x{"hello", "ok"}
	//fmt.Println(gutil.Export(v))
	//g.Dump(m)
	//g.Dump(v)
	//g.Client()

}
