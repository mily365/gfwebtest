
package service

import (
	"gfwebtest/app"

	"gfwebtest/app/service/base"
	"github.com/gogf/gf/util/gmeta"
)
var(
	UserSve *userSve
)
type userSve struct {
	gmeta.Meta `path:"service.user"`
	base.ServiceBase
}
func init() {
	UserSve=&userSve{gmeta.Meta{},base.ServiceBase{}}
	app.AppContext.RegisterObj(UserSve)
}

