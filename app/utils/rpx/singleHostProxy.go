package rpx

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type singleHostProxy struct {
	*httputil.ReverseProxy
}

func GetRandServer(hostCurrent string) interface{} {
	var findHost g.Map
	jUpstreams := g.Config().GetMap("upstreams")
	hosts := jUpstreams["hosts"]
	for _, v := range hosts.([]interface{}) {
		log.Println(v.(g.Map)["host"])
		log.Println(v.(g.Map)["backends"])
		if v.(g.Map)["host"] == hostCurrent {
			findHost = v.(g.Map)
			break
		}
	}
	//n := time.Now().Unix() % 2
	backendsLength := len(findHost["backends"].([]interface{}))
	n := time.Now().Unix() % int64(backendsLength)

	return findHost["backends"].([]interface{})[n]
}
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}
func NewSingleHostProxy2(targetHost string) (*singleHostProxy, error) {
	g.Dump("delegateToHost")
	director := func(req *http.Request) {
		delegateToHost := GetRandServer(targetHost).(g.Map)["ip"].(string)
		g.Dump(delegateToHost)
		urlToDelegate, err := url.Parse("http://" + delegateToHost)
		if err != nil {
			panic(err.Error())
		}
		req.URL.Scheme = urlToDelegate.Scheme
		req.URL.Host = urlToDelegate.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(urlToDelegate, req.URL)
		targetQuery := urlToDelegate.RawQuery
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	proxy := &httputil.ReverseProxy{Director: director}
	return &singleHostProxy{proxy}, nil
}

func NewSingleHostProxy(targetHost string) (*httputil.ReverseProxy, error) {

	urlObj, err := url.Parse(targetHost)

	if err != nil {

		return nil, err

	}

	proxy := httputil.NewSingleHostReverseProxy(urlObj)

	g.Dump(proxy)

	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {

		originalDirector(req)

		modifyRequest(req)

	}

	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil

}
func modifyRequest(req *http.Request) {
	req.Header.Set("Request-Id", "xxoo1122dddd")
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")

}
func modifyResponse() func(*http.Response) error {

	return func(resp *http.Response) error {

		resp.Header.Set("X-Proxy", "Magical")

		return nil

	}
}
func errorHandler() func(http.ResponseWriter, *http.Request, error) {

	return func(w http.ResponseWriter, req *http.Request, err error) {

		fmt.Printf("Got error while modifying response: %v \n", err)

		return

	}
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//从配置里面获取
		proxy.ServeHTTP(w, r)
	}

}
