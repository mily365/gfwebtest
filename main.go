package main

import (
	"log"
	"net/http"
	"xpass/app/utils/rpx"
)

func main() {

	// initialize a reverse proxy and pass the actual backend server url here

	// 初始化反向代理并传入真正后端服务的地址

	//proxy, err := rpx.NewSingleHostProxy("http://192.168.4.1:8190")
	//
	//if err != nil {
	//
	//	panic(err)
	//
	//}
	//
	//http.HandleFunc("/", rpx.ProxyRequestHandler(proxy))
	//
	//// handle all requests to your server using the proxy
	//// 使用 proxy 处理所有请求到你的服务
	//
	mhp := rpx.NewMultipleHostProxy()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//基本认证支持
		//auth := request.Header.Get("Authorization")
		//fmt.Println(auth)
		//if auth == "" {
		//	writer.Header().Set("WWW-Authenticate", `Basic realm="您必须输入用户名和密码"`)
		//	writer.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		mhp.ServeHTTP(writer, request)
	})

	//
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
	//log.Print(rpx.GetRandServer("192.168.4.1:8080"))

}
