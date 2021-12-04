package model

import (
	"github.com/gogf/gf/frame/g"
	"xpass/app"
)

type Search struct {
	//Test       string      `json:"test" v:"required#请输入用户姓名"`
	Tbl       string `json:"tbl"`
	Fields    string `json:"fields"`
	PageNo    uint   `json:"pageNo"`    // 用户ID
	PageSize  uint   `json:"PageSize"`  // 用户账号
	QueryForm *g.Map `json:"queryForm"` // 用户密码
}

func NewSearch() interface{} {
	var search *Search
	return &search
}
func init() {
	fun := NewSearch
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("search", fun)
}

type SearchResult struct {
	//Test       string      `json:"test" v:"required#请输入用户姓名"`
	Total int         `json:"total"`
	Rows  interface{} `json:"rows"`
}
