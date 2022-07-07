package main

import (
	"os"
)

type x struct {
	M string `json:"m"`
	N string `json:"n"`
}
type FieldType string

const (
	Id       FieldType = "id"
	Int      FieldType = "number"
	Text     FieldType = "text"
	DateTime FieldType = "datetime"
)

var tmpMap = map[string]interface{}{
	"text": "ddd",
}

func makeSqlOneStrByFieldType(propName string, fieldType FieldType) string {
	rtn := ""
	switch fieldType {
	case Id:
		rtn = `${FieldName} int(10) unsigned NOT NULL AUTO_INCREMENT`
	case DateTime:
		rtn = `${FieldName} datetime DEFAULT NULL`
	default:
		rtn = `${FieldName} varchar(${length}) NOT NULL,`
	}
	os.Expand(rtn, func(s string) string {
		return "xxx"
	})
	return "rtn"
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
