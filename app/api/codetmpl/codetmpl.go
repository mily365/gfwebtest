package codetmpl

var(
	ApiTmpl=`
package api

import (
	"gfwebtest/app"
	"gfwebtest/app/api/base"
	"github.com/gogf/gf/util/gmeta"
)

type {{.modelName}}Api struct {
	gmeta.Meta {{.path}}
	base.ApiBase
}
var (
	{{.exportModelName}}Api *{{.modelName}}Api
)
func init()  {
	{{.exportModelName}}Api=&{{.modelName}}Api{gmeta.Meta{},base.ApiBase{}}
	app.AppContext.RegisterObj({{.exportModelName}}Api)
}
`
	ModelTmpl=`
package model
import "gfwebtest/app"
func New{{.exportModelName}}() interface{}{
	var {{.modelName}} *{{.exportModelName}}
	return &{{.modelName}}
}
func New{{.exportModelName}}s() interface{}{
	var {{.modelName}}s []*{{.exportModelName}}
	return &{{.modelName}}s
}
func init() {
	fun:=New{{.exportModelName}}
	funs:=New{{.exportModelName}}s
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("{{.modelName}}", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("{{.modelName}}s", funs)
}
`
	ServiceTmpl=`
package service

import (
	"gfwebtest/app"

	"gfwebtest/app/service/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	{{.exportModelName}}Sve *{{.modelName}}Sve
)
type {{.modelName}}Sve struct {
	gmeta.Meta {{.path}}
	base.ServiceBase
}
func init() {
	{{.exportModelName}}Sve=&{{.modelName}}Sve{gmeta.Meta{},base.ServiceBase{}}
	app.AppContext.RegisterObj({{.exportModelName}}Sve)
}

`
	DaoTmpl=`
package dao

import (
	"gfwebtest/app"
	"gfwebtest/app/dao/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	{{.exportModelName}}Dao *{{.modelName}}Dao
)
type {{.modelName}}Dao struct {
	gmeta.Meta {{.path}}
	base.DaoBase
}
func init() {
	{{.exportModelName}}Dao=&{{.modelName}}Dao{gmeta.Meta{},base.DaoBase{}}
	app.AppContext.RegisterObj({{.exportModelName}}Dao)
}
`
)
