package codetmpl

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
	"xpass/app/model"
)

type initDataApi struct {
	gmeta.Meta `path:"api.initdata"`
	base.ApiBase
}

var (
	InitDataApi *initDataApi
)

func init() {
	InitDataApi = &initDataApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(InitDataApi)
}

func (cg *initDataApi) Initdata(r *ghttp.Request) {
	g.Log().Debug("xxxxxxxxxxxxxxxxxxxxxxxx")
	err, _ := app.ModelFactory.TxModelActions("solution", func(tx *gdb.TX, xmodel *gdb.Model) (error, interface{}) {
		controlInfoMetaData := model.Solution{}
		controlInfoMetaData.BizCode = "xpass_controlinfo"
		controlInfoMetaData.ModelName = "ControlInfo"
		controlInfoMetaData.Title = "实体UI描述信息元数据"
		xmodel.TX(tx).Delete("id>0")
		app.ModelFactory.GetModel("control_info").TX(tx).Delete("id>0")

		sid, err := xmodel.TX(tx).Data(controlInfoMetaData).OmitEmpty().InsertAndGetId()
		if err != nil {
			return err, nil
		}
		lst := g.List{
			g.Map{
				"sid":             sid,
				"title":           "标题",
				"propName":        "title",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "属性名",
				"propName":        "propName",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [] }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "是否主键",
				"propName":        "isKey",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "slide_toggle_开关选择",
				"groupName":       "开关属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [] }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "是否隐藏",
				"propName":        "isHidden",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "slide_toggle_开关选择",
				"groupName":       "开关属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "图标",
				"propName":        "icon",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "顺序",
				"propName":        "order",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "number_数字输入",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "控件类型",
				"propName":        "controlType",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "search_single_select_搜索单选",
				"groupName":       "选择属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [],
						"options": [
						    { "key": "text", "displayText": "文本输入" },
							{ "key": "chiplist_input", "displayText": "标签输入" },
							{ "key": "number", "displayText": "数字输入" },
							{ "key": "slide_toggle", "displayText": "开关选择" },
							{ "key": "label", "displayText": "文本只读" },
							{ "key": "datetime", "displayText": "日期时间" },
							{ "key": "datetime_span", "displayText": "日期范围" },
							{ "key": "search_single_select", "displayText": "搜索单选" },
							{ "key": "search_multi_select", "displayText": "搜索多选" },
							{ "key": "templ", "displayText": "模板" },
							{ "key": "actions", "displayText": "动作列" },
							{ "key": "action_button", "displayText": "动作按钮" },
							{ "key": "childTable", "displayText": "子表" }
			           ],
						"refDicCode": "false", "refModel": "false" }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "验证器",
				"propName":        "validatorType",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "search_multi_select_搜索多选",
				"groupName":       "选择属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [],
						"options": [
			                { "key": "none", "displayText": "无" },
							{ "key": "required", "displayText": "必输" },
							{ "key": "minLength", "displayText": "最小长度" },
							{ "key": "max", "displayText": "最大值" },
							{ "key": "min", "displayText": "最小值" }
			           ],
						"refDicCode": false, "refModel": false }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "显示位置",
				"propName":        "controlPosition",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "search_single_select_搜索单选",
				"groupName":       "选择属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [],
						"options": [
				            { "key": "all", "displayText": "所有" },
							{ "key": "list", "displayText": "列表" },
							{ "key": "form", "displayText": "表单" },
							{ "key": "search", "displayText": "搜索"}
			           ],
						"refDicCode": false, "refModel": false }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "分组名",
				"propName":        "groupName",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},

			g.Map{
				"sid":             sid,
				"title":           "操作",
				"propName":        "actions",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "actions_动作列",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "添加",
				"propName":        "add",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "删除",
				"propName":        "delete",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "修改",
				"propName":        "edit",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "取消",
				"propName":        "cancel",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "form_表单",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "保存",
				"propName":        "save",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "form_表单",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
		}
		_, err2 := app.ModelFactory.GetModel("control_info").TX(tx).Data(lst).Save()
		if err2 != nil {
			return err2, nil
		}
		solutionInfoMetaData := model.Solution{}
		solutionInfoMetaData.BizCode = "xpass_solution"
		solutionInfoMetaData.ModelName = "Solution"
		solutionInfoMetaData.Title = "解决方案"

		sid, err = xmodel.TX(tx).Data(solutionInfoMetaData).OmitEmpty().InsertAndGetId()
		if err != nil {
			return err, nil
		}
		lst = g.List{
			g.Map{
				"sid":             sid,
				"title":           "业务编码",
				"propName":        "bizCode",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "基本信息",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "显示名称",
				"propName":        "title",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "基本信息",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [] }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "实体名称",
				"propName":        "modelName",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "text_文本输入",
				"groupName":       "基本信息",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [] }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "方案引用",
				"propName":        "refSolutions",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "childTable_子表",
				"groupName":       "元数据描述",
				"validatorType":   "none_无",
				"controlPosition": "form_表单",
				"extraInfo": `{
						"isSort": false, "dynamicValues": [],
					    "options":[
						  {"key":"xpass_controlinfo","displayText":"方案实体属性描述"}
						 ]
			     }`,
			},
			g.Map{
				"sid":             sid,
				"title":           "操作",
				"propName":        "actions",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "actions_动作列",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "all_所有",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "添加",
				"propName":        "add",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "删除",
				"propName":        "delete",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "修改",
				"propName":        "edit",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "list_列表",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "取消",
				"propName":        "cancel",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "form_表单",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
			g.Map{
				"sid":             sid,
				"title":           "保存",
				"propName":        "save",
				"icon":            "",
				"isKey":           false,
				"isHidden":        false,
				"order":           "0",
				"controlType":     "action_button_动作按钮",
				"groupName":       "文本属性",
				"validatorType":   "none_无",
				"controlPosition": "form_表单",
				"extraInfo": `{
						"isSort": false, "dynamicValues": []}`,
			},
		}
		_, err2 = app.ModelFactory.GetModel("control_info").TX(tx).Data(lst).Save()
		if err2 != nil {
			return err2, nil
		}
		return nil, nil
	})
	if err != nil {
		panic(err.Error())
	}

	app.WrapSuccessRtn("init data ok", "ok", r)
}

//err, _ := app.ModelFactory.TxModelActions("user", func(tx *gdb.TX, xmodel *gdb.Model) (error, interface{}) {
//	//for i := 1; i <= 5; i++ {
//	//	// User.
//	//	user := model.User{
//	//		Name: fmt.Sprintf(`name_%d`, i),
//	//	}
//	//	lastInsertId, err := xmodel.TX(tx).Data(user).OmitEmpty().InsertAndGetId()
//	//	if err != nil {
//	//		return err, nil
//	//	}
//	//	// Detail.
//	//	userDetail := model.UserDetail{
//	//		Uid:     uint(lastInsertId),
//	//		Address: fmt.Sprintf(`address_%d`, lastInsertId),
//	//	}
//	//	_, err = app.ModelFactory.GetModel("user_detail").TX(tx).Data(userDetail).OmitEmpty().Insert()
//	//	if err != nil {
//	//		return err, nil
//	//	}
//	//	// Scores.
//	//	for j := 1; j <= 5; j++ {
//	//		userScore := model.UserScore{
//	//			Uid:    uint(lastInsertId),
//	//			Score:  uint(j),
//	//			Course: string(j),
//	//		}
//	//		_, err = app.ModelFactory.GetModel("user_score").TX(tx).Data(userScore).OmitEmpty().Insert()
//	//		if err != nil {
//	//			return err, nil
//	//		}
//	//	}
//	//}
//	//return nil, nil
//	//创建解决方案实体的元数据
//	solutonMetaData := model.Solution{}
//	solutonMetaData.BizCode = "xpass_solution"
//	solutonMetaData.ModelName = "Solution"
//	solutonMetaData.Title = "解决方案"
//
//	//业务编码	显示名称	实体名称	操作
//	//解决方案
//	//主表单信息
//	//方案实体属性描述
//	//[
//	//  { "title": "xsd",
//	//    "propName": "dfsd",
//	//    "icon": "solution",
//	//    "order": "0",
//	//    "groupName": "hello",
//	//    "version": 0,
//	//    "id": -1,
//	//    "isKey": true,
//	//    "isHidden": false,
//	//    "controlType": "search_multi_select_搜索多选",
//	//    "validatorType": "none_无",
//	//    "controlPostion": "all_所有",
//	//    "extraInfo": {
//	//         "isSort": false, "dynamicValues": [],
//	//         "options": [ { "key": "ok", "displayText": "好的" } ],
//	//         "refDicCode": false, "refModel": false }
//	//    }
//	//]
//
//	return nil, nil
//})
//if err != nil {
//	panic(err.Error())
//}
