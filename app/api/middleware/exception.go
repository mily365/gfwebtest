package middleware

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"xpass/app"

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
	//记录请求参数和访问的地址toes，进入es
	r.SetCtxVar(app.ResponseTimeKey, gtime.TimestampMilli())
	app.LoggerWithCtx(r.Context()).Async(true).Info(g.Map{"accsessPath": r.URL.Path, "reqParams": r.GetRequestMap()})
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件,写入ES
		app.LoggerWithCtx(r.Context()).Async(true).Error(gerror.Stack(err))

		//返回固定的友好信息
		r.Response.ClearBuffer()
		ctxInfo := r.GetCtxVar(app.ContextInfoKey).Interface().(*app.ContextInfo)
		r.Response.WriteJsonExit(ctxInfo.RtnInfo)
	} else {
		if gstr.Contains(r.URL.Path, "api") {
			ctxInfo := r.GetCtxVar(app.ContextInfoKey).Interface().(*app.ContextInfo)
			//成功执行操作后，记录审核日志到es

			app.LoggerWithCtx(r.Context()).Async(true).Info(g.Map{"accsessPath": r.URL.Path, "reqParams": r.GetRequestMap(), "resResult": ctxInfo.RtnInfo})
			r.Response.WriteJsonExit(ctxInfo.RtnInfo)

		}
	}
}
