
package api

import (
	"gfwebtest/app"
	"gfwebtest/app/api/base"
	"github.com/gogf/gf/util/gmeta"
)

type userApi struct {
	gmeta.Meta `path:"api.user"`
	base.ApiBase
}
var (
	UserApi *userApi
)
func init()  {
	UserApi=&userApi{gmeta.Meta{},base.ApiBase{}}
	app.AppContext.RegisterObj(UserApi)
}
