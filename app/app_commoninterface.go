package app

import "github.com/gogf/gf/net/ghttp"
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
}

// CommonOperation all service implement this interface
type CommonOperation interface {
	All(context.Context, interface{}) interface{}
	Withalls(context.Context, interface{}) interface{}
	Scrollpage(context.Context, interface{}) interface{}
	Create(context.Context, interface{}) interface{}
	Update(context.Context, interface{}) interface{}
	Delete(context.Context, interface{}) interface{}
}
