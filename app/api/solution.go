package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
)

type SolutionApi struct {
	gmeta.Meta `path:"api.solution"`
	base.ApiBase
}

var (
	solutionApi *SolutionApi
)

func init() {
	solutionApi = &SolutionApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(solutionApi)
}
func (p *SolutionApi) Initaddform(r *ghttp.Request) {
	g.Log().Debug("InitAddForm ..child..")
	_ = r.GetRequestMap()
	app.WrapSuccessRtn(nil, "ok", r)

}
