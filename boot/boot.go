package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"xpass/app"
	_ "xpass/app/dao"
	_ "xpass/app/dao/base"
	_ "xpass/app/dao/driver"
	_ "xpass/app/service"
	_ "xpass/app/service/base"
	_ "xpass/packed"
)

func init() {
	path2ModelMap := g.Config().GetMap("path2Model")
	model2TblMap := g.Config().GetMap("model2Tbl")
	//读取实体表类型的方案，构造配置对像需要的映射
	um := app.ModelFactory.GetModel("xpass_solution")
	res, _ := um.All(g.Map{"solution_type": "entityTable_实体表"})
	for _, v := range res.List() {
		modelName := v["model_name"].(string)
		tblName := v["biz_code"].(string)
		path := gstr.ToLower(modelName)
		path2ModelMap[path] = modelName
		model2TblMap[modelName] = tblName
	}
	path2ModelMap["solutionwithcontrolinfo"] = "SolutionWithControlInfo"
	path2ModelMap["codegen"] = "CodeGen"
	path2ModelMap["initdata"] = "InitData"
	g.Dump(path2ModelMap)
	g.Dump(model2TblMap)

}
