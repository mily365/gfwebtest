package router

import (
	"gfwebtest/app"
	_ "gfwebtest/app/api/base"
	_ "gfwebtest/app/api/codetmpl"

	_ "gfwebtest/app/api"
	_ "gfwebtest/app/api/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"reflect"
)

func init() {
	s := g.Server()
	for _, v := range app.AppContext.Keys() {
		//bind middlewares
		if gstr.HasPrefix(v, "middle") {
			strarray := gstr.Split(v, ".")
			mstr := "/" + gstr.Join(strarray[1:], "/")
			midobjValue := reflect.ValueOf(app.AppContext.Get(v))
			nummds := midobjValue.NumMethod()
			var funcs []func(*ghttp.Request)
			for n := 0; n < nummds; n++ {
				//隐式类型转换为eface,valueOf-iface-->任何一个成员都会转换为eface,需要进行断言类型转换
				funcs = append(funcs, midobjValue.Method(n).Interface().(func(*ghttp.Request)))
			}
			s.BindMiddleware(mstr, funcs...)
		}
		//bind api-handler
		if gstr.HasPrefix(v, "api") {
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
