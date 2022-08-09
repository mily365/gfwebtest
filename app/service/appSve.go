package service

import (
	"context"
	"github.com/gogf/gf/util/guid"
	"xpass/app"
	"xpass/app/model"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	AppSve *appSve
)

type appSve struct {
	gmeta.Meta `path:"service.app"`
	base.ServiceBase
}

func init() {
	AppSve = &appSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(AppSve)
}
func (s *appSve) InitAddForm(ctx context.Context, i interface{}) interface{} {
	appObj := new(model.App)
	appObj.AppKey = guid.S()
	return appObj
}
