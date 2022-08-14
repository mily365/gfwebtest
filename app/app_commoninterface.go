package app

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/net/ghttp"
)
import "context"

// ApiOperation all api implement  this interface
type ApiOperation interface {
	InitAddForm(r *ghttp.Request)
	All(r *ghttp.Request)
	Withalls(r *ghttp.Request)
	Scrollpage(r *ghttp.Request)
	Create(r *ghttp.Request)
	Update(r *ghttp.Request)
	Delete(r *ghttp.Request)
	Createtx(r *ghttp.Request)
	Updatetx(r *ghttp.Request)
	Deletetx(r *ghttp.Request)
	Copytx(r *ghttp.Request)
	Findone(r *ghttp.Request)
}
type ServiceOperation interface {
	FindOne(context.Context, interface{}) interface{}
	InitAddForm(context.Context, interface{}) interface{}
	All(context.Context, interface{}) interface{}
	Withalls(context.Context, interface{}) interface{}
	Scrollpage(context.Context, interface{}) interface{}
	Create(context.Context, interface{}) interface{}
	Update(context.Context, interface{}) interface{}
	Delete(context.Context, interface{}) interface{}
	Createtx(context.Context, interface{}) interface{}
	Updatetx(context.Context, interface{}) interface{}
	Deletetx(context.Context, interface{}) interface{}
	Copytx(context.Context, interface{}) interface{}
}

// CommonOperation all service implement this interface
type DaoOperation interface {
	FindOne(context.Context, interface{}) interface{}
	All(context.Context, interface{}) interface{}
	Withalls(context.Context, interface{}) interface{}
	Scrollpage(context.Context, interface{}) interface{}
	Create(context.Context, interface{}) interface{}
	Update(context.Context, interface{}) interface{}
	Delete(context.Context, interface{}) interface{}
	Createtx(context.Context, interface{}, *gdb.TX, *gdb.Model) interface{}
	CreateEsTx(context.Context, interface{}, string)
	Updatetx(context.Context, interface{}, *gdb.TX, *gdb.Model) interface{}
	Deletetx(context.Context, interface{}, *gdb.TX, *gdb.Model) interface{}
	Copytx(context.Context, interface{}, *gdb.TX, *gdb.Model) interface{}
}
