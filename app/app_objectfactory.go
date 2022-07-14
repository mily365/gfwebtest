package app

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmeta"
	"reflect"
)

var (
	ApiPathPrefix      = "api"
	ServicePathPrefix  = "service"
	DaoPathPrefix      = "dao"
	ApiAdapterPath     = "api.*"
	ServiceAdapterPath = "service.*"
	DaoAdapterPath     = "dao.*"
	AopPrefix          = "middle"
)

type objectFactory struct {
	*gmap.StrAnyMap
}

func (of *objectFactory) GetObject(path string) interface{} {
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
