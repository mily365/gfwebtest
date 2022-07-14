package base

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"xpass/app"
)

type ApiBase struct {
	Sve app.ServiceOperation
}

func (p *ApiBase) Scrollpage(r *ghttp.Request) {
	g.Log().Debug("Scrollpage alls....")
	q := r.GetRequestMap()
	s := p.Sve.Scrollpage(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) Initaddform(r *ghttp.Request) {
	g.Log().Debug("InitAddForm ....")
	_ = r.GetRequestMap()
	app.WrapSuccessRtn(nil, "ok", r)

}

func (p *ApiBase) Withalls(r *ghttp.Request) {
	g.Log().Debug("Withalls alls....")
	q := r.GetRequestMap()
	s := p.Sve.Withalls(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) All(r *ghttp.Request) {
	g.Log().Debug("ibase all....", r.Context().Value(app.Path2ModelRegKey))
	q := r.GetRequestMap()
	s := p.Sve.All(r.Context(), q)
	app.WrapSuccessRtn(s, "ok", r)

}

func (p *ApiBase) Create(r *ghttp.Request) {
	modelKey := r.GetCtxVar(app.Path2ModelRegKey).String()
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
func (p *ApiBase) Createtx(r *ghttp.Request) {
	g.Log().Debug("Createtx")
	modelKey := r.GetCtxVar(app.Path2ModelRegKey).String()
	typeStruct := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	r.Parse(typeStruct)
	err := g.Validator().CheckStruct(typeStruct)
	g.Dump(typeStruct)
	if err != nil {
		panic(err.Error())
	}
	toCreate := r.GetRequestMap()
	rtn := p.Sve.Createtx(r.Context(), toCreate)
	app.WrapSuccessRtn(rtn, "ok", r)
}

func (p *ApiBase) Update(r *ghttp.Request) {
	modelKey := r.GetCtxVar(app.Path2ModelRegKey).String()
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
func (p *ApiBase) Updatetx(r *ghttp.Request) {
	modelKey := r.GetCtxVar(app.Path2ModelRegKey).String()
	typeStruct := app.TypePointerFuncFactory.GetStructPointer(modelKey)
	r.Parse(typeStruct)
	err := g.Validator().CheckStruct(typeStruct)
	g.Dump(typeStruct)
	if err != nil {
		panic(err.Error())
	}
	toUpdate := r.GetRequestMap()
	rtn := p.Sve.Updatetx(r.Context(), toUpdate)
	g.Dump(rtn)
	//如果更新返回有值，那么判断为并发
	if rtn == nil {
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
