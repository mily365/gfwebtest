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
	"xpass/app"
)

type singleHostProxy struct {
	*httputil.ReverseProxy
}

func GetRandServer(hostCurrent string) interface{} {
	var findHost g.Map
	jUpstreams := g.Config("upstreams").GetMap("upstreams")
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

type director func(req *http.Request)

func NewSingleHostProxy2() (*singleHostProxy, error) {
	directorFun := func(req *http.Request) {
		//到达代理的来访目标req.Host,根据配置获取对应的要分发到的服务器
		//根据请求的url  /serviceName/api/entity/aciotn
		//提取serviceName,按照serviceName找到具体的被委托主机
		//被委托的服务启动时，要进行服务注册，注册的路径 /services/student-service/
		//注意req是被科隆后的
		delegateToHost := GetRandServer(req.Host).(g.Map)["ip"].(string)

		app.Logger.Warning("current delegate ip is " + delegateToHost)
		urlToDelegate, err := url.Parse("http://" + delegateToHost)
		if err != nil {
			panic(err.Error())
		}
		req.URL.Scheme = urlToDelegate.Scheme
		req.URL.Host = urlToDelegate.Host

		app.Logger.Warning(req.URL)

		req.URL.Path, req.URL.RawPath = joinURLPath(urlToDelegate, req.URL)

		app.Logger.Warning(req.URL.Path)
		app.Logger.Warning(req.URL.RawPath)
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
	app.Logger.Warning("ppppppp")
	proxy := &httputil.ReverseProxy{Director: directorFun}

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)

	}

	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
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
	req.Header.Set("X-Forwarded-For-Host", req.Host)
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")

	for r, v := range req.Header {
		app.Logger.Warning(r, v)
	}
	jiaoyiToken := "Bearer jiaoyi"
	smToken := "Bearer sm"
	zzbToken := "Bearer jytest6"
	if req.Header["Authorization"] != nil {
		if smToken == req.Header["Authorization"][0] {
			app.Logger.Warning(req.Header["Authorization"][0])
			req.Header.Set("x-consumer-username", "sm")
			req.Header.Set("x-consumer-custom-id", "1")
			req.Header.Set("x-credential-identifier", "Simple-Reverse-Proxy")
			req.Header.Set("x-consumetag", "cmp_1|pass_3b56ba476dc4f75603a59e06ed470406")
		}
		if jiaoyiToken == req.Header["Authorization"][0] {
			app.Logger.Warning(req.Header["Authorization"][0])
			req.Header.Set("x-consumer-username", "jiaoyi")
			req.Header.Set("x-consumer-custom-id", "60")
			req.Header.Set("x-credential-identifier", "Simple-Reverse-Proxy")
			req.Header.Set("x-consumetag", "cmp_40|pass_516f23ffe1eb6628f2898c919b1d9fad")

		}
		if zzbToken == req.Header["Authorization"][0] {
			app.Logger.Warning(req.Header["Authorization"][0])
			req.Header.Set("x-consumer-username", "jytest6")
			req.Header.Set("x-consumer-custom-id", "10")
			req.Header.Set("x-credential-identifier", "Simple-Reverse-Proxy")
			req.Header.Set("x-consumetag", "cmp_10|pass_7c8712d838190c3fdc6305aa7584bc98")

		}

	}

	//如果host是交付中心平台那么，就设置超管身份
	//if gstr.Equal(req.Host, "t9.com:8000") {
	//	req.Header.Set("x-consumer-username", "sm")
	//	req.Header.Set("x-consumer-custom-id", "1")
	//	req.Header.Set("x-credential-identifier", "Simple-Reverse-Proxy")
	//	req.Header.Set("x-consumetag", "cmp_1|pass_3b56ba476dc4f75603a59e06ed470406")
	//}
	//req.Header.Set("x-consumer-username", "Simple-Reverse-Proxy")
	//req.Header.Set("x-consumer-custom-id", "Simple-Reverse-Proxy")
	//req.Header.Set("x-credential-identifier", "Simple-Reverse-Proxy")
	//req.Header.Set("x-consumetag", "cmp_1|pass_3b56ba476dc4f75603a59e06ed470406")

}
func modifyResponse() func(*http.Response) error {

	return func(resp *http.Response) error {
		app.Logger.Warning("resp....................................")
		resp.Header.Add("Access-Control-Allow-Origin", "*")
		resp.Header.Add("Access-Control-Allow-Headers", "*")
		////resp.Header.Add("Access-Control-Allow-Credentials", "true")
		resp.Header.Add("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
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
