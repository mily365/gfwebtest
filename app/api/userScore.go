
package api

import (
	"gfwebtest/app"
	"gfwebtest/app/api/base"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
)

type userScoreApi struct {
	gmeta.Meta `path:"api.userScore"`
	base.ApiBase
}
var (
	UserScoreApi *userScoreApi
)
func init()  {
	UserScoreApi=&userScoreApi{gmeta.Meta{},base.ApiBase{}}
	app.AppContext.RegisterObj(UserScoreApi)
}
func (p *userScoreApi) All(r *ghttp.Request){
	g.Log().Debug("UserScoreApi all....",r.Context().Value("tbl"))
	q:=r.GetRequestMap()
	s:=p.Sve.(app.CommonOperation).All(r.Context(),q)
	app.WrapSuccessRtn(s,"ok",r)
}