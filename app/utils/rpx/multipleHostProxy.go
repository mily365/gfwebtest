package rpx

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"net/http"
	"xpass/app"
)

type multipleHostProxy struct {
	*gmap.StrAnyMap
}

func NewMultipleHostProxy() *multipleHostProxy {
	return &multipleHostProxy{gmap.NewStrAnyMap(true)}
}
func (mhp *multipleHostProxy) GetOrSetSingleHostProxy(hostName string) *singleHostProxy {
	proxy, err := NewSingleHostProxy2()
	if err != nil {
		panic(err.Error())
	}
	return mhp.GetOrSet(hostName, proxy).(*singleHostProxy)
}

func (mhp *multipleHostProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Logger.Debug(r.Host)
	shp := mhp.GetOrSetSingleHostProxy(r.Host)
	g.Dump(shp)
	shp.ServeHTTP(w, r)
}
