package router

import (
	"xpass/app"
	_ "xpass/app/api/base"
	_ "xpass/app/api/codetmpl"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"reflect"
	_ "xpass/app/api"
	_ "xpass/app/api/middleware"
)

func init() {
	s := g.Server()
	for _, v := range app.AppContext.Keys() {
		//bind middlewares
		if gstr.HasPrefix(v, app.AopPrefix) {
			strarray := gstr.Split(v, ".")
			mstr := "/" + gstr.Join(strarray[1:], "/")
			app.Logger.Debug(mstr, "vvvvvvvvvvvvvvvvvvvvvv")
			midobjValue := reflect.ValueOf(app.AppContext.Get(v))
			nummds := midobjValue.NumMethod()
			var funcs []func(*ghttp.Request)
			for n := 0; n < nummds; n++ {
				//隐式类型转换为eface,valueOf-iface-->任何一个成员都会转换为eface,需要进行断言类型转换
				//值对象访问原始对象，Int()....
				funcs = append(funcs, midobjValue.Method(n).Interface().(func(*ghttp.Request)))
			}
			s.BindMiddleware(mstr, funcs...)
		}
		//bind api-handler
		if gstr.HasPrefix(v, app.ApiPathPrefix) {
			rstr := "/" + gstr.Join(gstr.Split(v, "."), "/")
			s.BindObject(rstr, app.AppContext.Get(v))
		}
		// bind exceptionHandler
		if gstr.HasPrefix(v, "exp") {
			midobjValue := reflect.ValueOf(app.AppContext.Get(v))
			s.BindMiddlewareDefault(midobjValue.Method(0).Interface().(func(*ghttp.Request)))
		}

	}
}
