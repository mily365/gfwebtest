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

func (*uploadSve) All(i interface{}) interface{} {
	app.Logger.Debug("adapterSve......All", i)
	return "all called"
}

func (*uploadSve) Create(i interface{}) interface{} {
	panic("implement me")
}

func (*uploadSve) Update(i interface{}) interface{} {
	panic("implement me")
}

func (*uploadSve) Delete(i interface{}) interface{} {
	panic("implement me")
}
