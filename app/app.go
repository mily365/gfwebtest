package app

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var (
	AppContext = objectFactory{
		gmap.NewStrAnyMap(true),
	}
	ModelFactory = modelFactory{
		gmap.NewStrAnyMap(true),
	}
	TypePointerFuncFactory = typePointerFuncFactory{
		make(map[string]func() interface{}),
	}
)

type AppError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Ext  interface{} `json:"ext"`
}

func (receiver *AppError) Error() string {
	return receiver.Msg
}
func (receiver *AppError) Extra() interface{} {
	return receiver.Ext
}

type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

//-------------------context info----------------------
const (
	ContextInfoKey   = "ContextInfoKey"
	SessionKeyUser   = "SessionKeyUser"
	TraceID          = "TraceID"
	Path2ModelRegKey = "PathModelName"
	ResponseTimeKey  = "ResponseTimeKey"
)

//config key
const (
	ModelToTbl = "model2Tbl"
	Path2Model = "path2Model"
)

type ContextInfo struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
	RtnInfo *RtnInfo
}

// ContextUser 请求上下文中的用户信息
type ContextUser struct {
	Id       uint   // 用户ID
	Name     string // 用户账号
	Nickname string // 用户名称
}

//--------rtnInfo-----
type RtnInfo struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func WrapSuccessRtn(d interface{}, msg string, r *ghttp.Request) RtnInfo {
	rtn := RtnInfo{
		Code: 0,
		Data: d,
		Msg:  msg,
	}
	ctxinfo := r.GetCtxVar(ContextInfoKey).Interface().(*ContextInfo)
	ctxinfo.RtnInfo = &rtn
	return rtn
}
func WrapFailRtn(d interface{}, msg string, r *ghttp.Request) RtnInfo {
	rtn := RtnInfo{
		Code: -1,
		Data: d,
		Msg:  msg,
	}
	ctxInfo := r.GetCtxVar(ContextInfoKey).Interface().(*ContextInfo)
	ctxInfo.RtnInfo = &rtn
	panic(d)
	return rtn
}

type SearchResult struct {
	//Test       string      `json:"test" v:"required#请输入用户姓名"`
	Total    int         `json:"total"`
	Rows     interface{} `json:"rows"`
	ScrollId string      `json:"scrollId"`
}

func PageParam(search g.Map) (int, int) {
	var no int
	var ps int
	if search["pageNo"] != nil {
		//no=search["pageNo"]
		no = gconv.Int(search["pageNo"].(json.Number).String())
		if no == 0 {
			no = 1
		}
	} else {
		no = 1
	}
	if search["pageSize"] != nil {
		ps = gconv.Int(search["pageSize"].(json.Number).String())
	} else {
		ps = 10
	}
	return (no - 1) * ps, ps
}

func GetEsName(modelName string) (string, string) {
	appName := g.Config().GetString("appInfo.name")
	mName := gstr.ToLower(modelName)
	esName := fmt.Sprintf("%s_%s", appName, mName)
	return esName, appName
}
