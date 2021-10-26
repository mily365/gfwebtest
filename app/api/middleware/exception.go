package middleware

import (
	"gfwebtest/app"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
)

var (
	Exp *exp
)

func init() {
	Exp = new(exp)
	app.AppContext.RegisterObj(Exp)
}

type exp struct {
	gmeta.Meta `path:"exp"`
}

func (*exp) ExceptionHandler(r *ghttp.Request) {

	r.Middleware.Next()

	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件,写入ES
		app.LoggerWithCtx(r.Context()).Error(gerror.Stack(err))
		//返回固定的友好信息
		r.Response.ClearBuffer()
		ctxInfo := r.GetCtxVar(app.ContextInfoKey).Interface().(*app.ContextInfo)
		r.Response.WriteJsonExit(ctxInfo.RtnInfo)
	} else {
		if gstr.Contains(r.URL.Path, "api") {
			ctxInfo := r.GetCtxVar(app.ContextInfoKey).Interface().(*app.ContextInfo)
			r.Response.WriteJsonExit(ctxInfo.RtnInfo)
		}
	}
}
