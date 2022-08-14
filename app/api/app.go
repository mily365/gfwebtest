package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
	base2 "xpass/app/service/base"
)

type appApi struct {
	gmeta.Meta `path:"api.app"`
	base.ApiBase
}

var (
	AppApi *appApi
)

func init() {
	AppApi = &appApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(AppApi)
}
func (application *appApi) Fetchapp(r *ghttp.Request) {
	g.Log().Debug("Fetchapp ....")
	q := r.GetRequestMap()
	rtn := app.AppContext.GetObject("service.app").(base2.AppServiceInterface).FetchApp(r.Context(), q)
	app.WrapSuccessRtn(rtn, "ok", r)
}
