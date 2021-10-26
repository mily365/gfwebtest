package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmeta"
	"reflect"
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

// ApiOperation all api implement  this interface
type ApiOperation interface {
	All(r *ghttp.Request)
	Withalls(r *ghttp.Request)
	Create(r *ghttp.Request)
	Update(r *ghttp.Request)
	Delete(r *ghttp.Request)
}

// CommonOperation all service implement this interface
type CommonOperation interface {
	All(context.Context, interface{}) interface{}
	Withalls(context.Context, interface{}) interface{}
	Create(context.Context, interface{}) interface{}
	Update(context.Context, interface{}) interface{}
	Delete(context.Context, interface{}) interface{}
}

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

// Logger ----------------------------------
var Logger = g.Log()

func LoggerWithCtx(ctx context.Context) *glog.Logger {
	Logger.SetHandlers(LoggingJsonHandler)
	return Logger.Ctx(ctx)
}

// JsonOutputsForLogger LoggingJsonHandler is a example handler for logging JSON format content.
type JsonOutputsForLogger struct {
	TraceId string `json:"traceId"`
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	if ctx.Value(TraceID) != nil {
		jsonForLogger := JsonOutputsForLogger{
			TraceId: ctx.Value(TraceID).(string),
			Time:    in.TimeFormat,
			Level:   gstr.Trim(in.LevelFormat, "[]"),
			Content: gstr.Trim(in.Content),
		}
		jsonBytes, err := json.Marshal(jsonForLogger)
		if err != nil {
			panic(errors.New("json log handler error!"))
		}
		if in.Level == glog.LEVEL_ERRO {
			// to es error
			fmt.Print("error.......")
		}
		if in.Level == glog.LEVEL_INFO {
			// to es
			fmt.Print("info.......")
		}
		in.Buffer().Write(jsonBytes)
		//in.Buffer().WriteString("\n")
		in.Content = string(jsonBytes)
	}

	//fmt.Println("to push es....")
	//fmt.Println(string(jsonBytes))

	//in.Content=string(jsonBytes)+"\n"
	//to do write to es......
	in.Next()

}

//----------------------------------------
type modelFactory struct {
	*gmap.StrAnyMap
}

func (mf *modelFactory) GetModel(tblName string) *gdb.Model {
	m := mf.GetOrSet(tblName, g.DB().Model(tblName)).(*gdb.Model)
	return m.Clone()
}
func (mf *modelFactory) TxModelActions(tblName string, fun func(tx *gdb.TX, model *gdb.Model) (error, interface{})) (error, interface{}) {
	tx, ex := g.DB().Begin()
	if ex != nil {
		panic(ex.Error())
	}
	m := mf.GetOrSet(tblName, g.DB().Model(tblName)).(*gdb.Model)
	err, rt := fun(tx, m.Clone())
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return err, rt
}

//------------------------
type typePointerFuncFactory struct {
	mapFuncPointer map[string]func() interface{}
}

func (tpf *typePointerFuncFactory) RegisterOrGetTypePointer(typeName string, typeFunPointer func() interface{}) func() interface{} {
	g.Log().Info(tpf.mapFuncPointer)
	if tpf.mapFuncPointer[typeName] != nil {
		return tpf.mapFuncPointer[typeName]
	} else {
		if typeFunPointer == nil {
			panic("please input param typeFunPointer...or check whether generate model register code!")
		}
		tpf.mapFuncPointer[typeName] = typeFunPointer
		return typeFunPointer
	}
}
func (tpf *typePointerFuncFactory) GetStructPointer(explicitName string) interface{} {
	fp := TypePointerFuncFactory.RegisterOrGetTypePointer(explicitName, nil)
	up := fp()
	return up
}
func (tpf *typePointerFuncFactory) GetStructArrayPointer(explicitName string) interface{} {
	fp := TypePointerFuncFactory.RegisterOrGetTypePointer(explicitName+"s", nil)
	up := fp()
	return up
}

//-----

type objectFactory struct {
	*gmap.StrAnyMap
}
type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

var (
	ApiPathPrefix      = "api"
	ServicePathPrefix  = "service"
	DaoPathPrefix      = "dao"
	ApiAdapterPath     = "api.*"
	ServiceAdapterPath = "service.*"
	DaoAdapterPath     = "dao.*"
)

type SetRef interface {
	SetRef(ref interface{})
}

func (of *objectFactory) getObject(path string) interface{} {
	return of.Get(path)
}
func (of *objectFactory) RegisterObj(child interface{}) {
	md := gmeta.Data(child)
	path := md["path"].(string)
	Logger.Debug(path, "component  register in appcontext")
	//inject
	//case api. inject service.
	if gstr.HasPrefix(path, ApiPathPrefix) {
		v := reflect.ValueOf(child).Elem().FieldByName("Sve")
		childTypeName := reflect.TypeOf(child).String()
		//apistruct ust contain apibase
		if !v.CanSet() {
			Logger.Warning("please defind Parent base.ApiBase in api struct!", path)
			panic("please defind Parent base.ApiBase in api struct!")
		} else {
			spath := gstr.Replace(path, ApiPathPrefix, ServicePathPrefix)
			daopath := gstr.Replace(path, ApiPathPrefix, DaoPathPrefix)
			serviceObj, isExist := of.Search(spath)

			if isExist {
				serviceTypeName := reflect.TypeOf(serviceObj).String()
				serviceValue := reflect.ValueOf(serviceObj)
				daoRef := serviceValue.Elem().FieldByName("Dao")
				daoObj, isDaoExist := of.Search(daopath)
				if isDaoExist {
					daoValue := reflect.ValueOf(daoObj)
					daoRef.Set(daoValue)
					Logger.Debug(path, "---", "inject", reflect.TypeOf(daoObj), "into", serviceTypeName)
				} else {
					daoAdapterObj := of.Get(DaoAdapterPath)
					daoAdapterValue := reflect.ValueOf(daoAdapterObj)
					daoRef.Set(daoAdapterValue)
					Logger.Debug(path, "---", "inject", reflect.TypeOf(daoAdapterObj), "into", serviceTypeName)
				}
				v.Set(serviceValue)
				Logger.Debug(path, "---", "inject", serviceTypeName, "into", childTypeName)

			} else {
				// not define service,use serviceAdapter
				serviceAdapaterObj := of.Get(ServiceAdapterPath)
				serviceValue := reflect.ValueOf(serviceAdapaterObj)
				v.Set(serviceValue)
				Logger.Warning(path, "---", " service not found..", "inject", reflect.TypeOf(serviceAdapaterObj), "into", reflect.TypeOf(child))
			}

		}
	}
	of.SetIfNotExist(path, child)
}

//-------------------context info----------------------
var (
	ContextInfoKey = "ContextInfoKey"
	SessionKeyUser = "SessionKeyUser"
	TraceID        = "TraceID"
	PathModelName  = "PathModelName"
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
