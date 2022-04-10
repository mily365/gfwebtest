package service

import (
	"xpass/app"

	"github.com/gogf/gf/util/gmeta"
	"xpass/app/service/base"
)

var (
	SolutionSve *solutionSve
)

type solutionSve struct {
	gmeta.Meta `path:"service.solution"`
	base.ServiceBase
}

func init() {
	SolutionSve = &solutionSve{gmeta.Meta{}, base.ServiceBase{}}
	app.AppContext.RegisterObj(SolutionSve)
}
