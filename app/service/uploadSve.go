package service

import (
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	UploadSve *uploadSve
)

func init() {
	UploadSve = &uploadSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(UploadSve)

}

type uploadSve struct {
	gmeta.Meta `path:"service.upload"`
	base.ServiceBase
}
