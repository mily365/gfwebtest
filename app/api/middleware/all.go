package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/guid"
	"xpass/app"
	"xpass/app/model"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
)

var (
	All *all
)

func init() {
	All = new(all)
	app.AppContext.RegisterObj(All)
}

type all struct {
	gmeta.Meta `path:"middle.*"`
}

func (*all) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (*all) Ctx(r *ghttp.Request) {

	if gstr.Contains(r.URL.Path, "api") {
		modelNameLower := gstr.Split(r.URL.Path, "/")[2]
		modelNme := g.Config().Get("path2Model." + gstr.ToLower(modelNameLower))
		if modelNme == nil {
			panic("access url not exist!")
		}
		////model type reg key
		r.SetCtxVar(app.PathModelName, gstr.CaseCamelLower(modelNme.(string)))

	}
	//set traceid
	r.SetCtxVar(app.TraceID, guid.S())
	contextInfo := &app.ContextInfo{}
	contextInfo.Session = r.Session
	if v := r.Session.GetVar(app.SessionKeyUser); !v.IsNil() {
		var user *model.User
		_ = v.Struct(&user)
		contextInfo.User = &app.ContextUser{
			Id:   user.Id,
			Name: user.Name,
		}
	}
	r.SetCtxVar(app.ContextInfoKey, contextInfo)

	r.Middleware.Next()
}
