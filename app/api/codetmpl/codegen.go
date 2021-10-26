package codetmpl

import (
	"context"
	"gfwebtest/app"
	"gfwebtest/app/api/base"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmeta"
)

type codeGenApi struct {
	gmeta.Meta `path:"api.codegen"`
	base.ApiBase
}

var (
	CodeGenApi *codeGenApi
)

func init() {
	CodeGenApi = &codeGenApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(CodeGenApi)
}

func (cg *codeGenApi) Gmodel(r *ghttp.Request) {
	pm := r.GetRequestMap()
	if pm["model"] == nil {
		panic("input model param....")
	}
	if pm["target"] == nil {
		panic("input target param....")
	}
	modelName := pm["model"].(string)

	exportModelName := modelName
	modelName = gstr.CaseCamelLower(modelName)

	target := pm["target"].(string)

	apiPath := g.Config().GetMap("codegen")["apiPath"].(string)
	modelPath := g.Config().GetMap("codegen")["modelPath"].(string)
	servicePath := g.Config().GetMap("codegen")["servicePath"].(string)
	daoPath := g.Config().GetMap("codegen")["daoPath"].(string)
	apifpath := apiPath + modelName + ".go"
	modelfpath := modelPath + modelName + ".go"
	servicefpath := servicePath + modelName + "Sve.go"
	daofpath := daoPath + modelName + "Dao.go"
	switch target {
	case "api":
		content, err := g.View().ParseContent(context.TODO(), ApiTmpl, g.Map{
			"path":            "`path:\"api." + modelName + "\"`",
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		gfile.PutContents(apifpath, content)
	case "model":

		content, err := g.View().ParseContent(context.TODO(), ModelTmpl, g.Map{
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		err = gfile.PutContents(modelfpath, content)
		g.Log().Debug(modelfpath)
		if err != nil {
			g.Log().Debug(err.Error())
		}

	case "service":
		content, err := g.View().ParseContent(context.TODO(), ServiceTmpl, g.Map{
			"path":            "`path:\"service." + modelName + "\"`",
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		gfile.PutContents(servicefpath, content)

	case "dao":
		content, err := g.View().ParseContent(context.TODO(), DaoTmpl, g.Map{
			"path":            "`path:\"dao." + modelName + "\"`",
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		gfile.PutContents(daofpath, content)

	default:
		g.Log().Debug("will add new target......")

	}
	app.WrapSuccessRtn("build"+target+"success.....", "ok", r)
}
