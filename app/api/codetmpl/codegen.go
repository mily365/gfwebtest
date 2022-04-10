package codetmpl

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmeta"
	"github.com/gogf/gf/util/gutil"
	"reflect"
	"xpass/app"
	"xpass/app/api/base"
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

//根据实体ES索引
func (cg *codeGenApi) Initesmaping(r *ghttp.Request) {
	mapFuncModelPointer := app.TypePointerFuncFactory.GetFuncMapForModelPointer()
	esNames := make([]string, 0)
	for _, mfp := range mapFuncModelPointer {
		modelPointer := mfp()
		tpy := reflect.TypeOf(modelPointer).Elem().Elem()
		tpyStr := tpy.Kind().String()
		if gstr.Equal(tpyStr, "struct") {
			eName := app.GetEsFactory().CreateIndex(r.Context(), tpy)
			esNames = append(esNames, eName)
		}
	}
	app.WrapSuccessRtn(esNames, "索引创建成功!", r)
}

func (cg *codeGenApi) Pinges(r *ghttp.Request) {
	g.Log().Debug("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	niF, _ := app.GetEsFactory().Client.NodesInfo().Do(r.Context())
	g.Dump(niF)
	app.WrapSuccessRtn(niF, "ok", r)
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
	//检查是否存在
	modelNameConfig := g.Cfg().GetString("path2Model." + gstr.ToLower(modelName))
	if gutil.IsEmpty(modelNameConfig) {
		panic("please config entity name " + modelName + " in config file...........................!")
	}

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
		if gfile.Exists(apifpath) == false {
			err = gfile.PutContents(apifpath, content)
			if err != nil {
				g.Log().Debug(err.Error())
			}

		}

	case "model":

		content, err := g.View().ParseContent(context.TODO(), ModelTmpl, g.Map{
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		//
		if gfile.Exists(modelfpath) == false {
			err = gfile.PutContents(modelfpath, content)
			g.Log().Debug(modelfpath)
			if err != nil {
				g.Log().Debug(err.Error())
			}
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
		if gfile.Exists(servicefpath) == false {
			err = gfile.PutContents(servicefpath, content)
			if err != nil {
				g.Log().Debug(err.Error())
			}
		}

	case "dao":
		content, err := g.View().ParseContent(context.TODO(), DaoTmpl, g.Map{
			"path":            "`path:\"dao." + modelName + "\"`",
			"modelName":       modelName,
			"exportModelName": exportModelName,
		})
		if err != nil {
			panic(err)
		}
		if gfile.Exists(daofpath) == false {
			err = gfile.PutContents(daofpath, content)
			if err != nil {
				g.Log().Debug(err.Error())
			}
		}

	default:
		g.Log().Debug("will add new target......")

	}
	app.WrapSuccessRtn("build"+target+"success.....", "ok", r)
}
