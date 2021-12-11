package base

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"xpass/app"
)

type ApiBase struct {
	Sve app.CommonOperation
}

func (p *ApiBase) Scrollpage(r *ghttp.Request) {
	g.Log().Debug("Scrollpage alls....")
	q := r.GetRequestMap()
	s := p.Sve.Scrollpage(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) Withalls(r *ghttp.Request) {
	g.Log().Debug("Withalls alls....")
	q := r.GetRequestMap()
	s := p.Sve.Withalls(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) All(r *ghttp.Request) {
	g.Log().Debug("ibase all....", r.Context().Value("tbl"))
	q := r.GetRequestMap()
	s := p.Sve.All(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) Create(r *ghttp.Request) {
	modelKey := r.GetCtxVar(app.PathModelName).String()
	g.Log().Debug(app.PathModelName)
	g.Log().Debug(modelKey)
	typeStruct := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	r.Parse(typeStruct)
	err := g.Validator().CheckStruct(typeStruct)
	g.Dump(typeStruct)
	if err != nil {
		panic(err.Error())
	}
	toCreate := r.GetRequestMap()
	rtn := p.Sve.Create(r.Context(), toCreate)
	app.WrapSuccessRtn(rtn, "ok", r)
}

func (p *ApiBase) Update(r *ghttp.Request) {
	modelKey := r.GetCtxVar(app.PathModelName).String()
	typeStruct := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	r.Parse(typeStruct)
	err := g.Validator().CheckStruct(typeStruct)
	g.Dump(typeStruct)
	if err != nil {
		panic(err.Error())
	}
	toUpdate := r.GetRequestMap()
	rtn := p.Sve.Update(r.Context(), toUpdate)
	if rtn != nil {
		appE := app.AppError{Msg: "update concurrent.....", Code: 1, Ext: rtn}
		app.WrapFailRtn(appE, "has error!", r)

	} else {

		app.WrapSuccessRtn(rtn, "call success!", r)
	}

}

func (p *ApiBase) Delete(r *ghttp.Request) {
	toDelete := r.GetRequestMap()
	rtn := p.Sve.Delete(r.Context(), toDelete)
	app.WrapSuccessRtn(rtn, "ok", r)
}
