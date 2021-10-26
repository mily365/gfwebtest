package main

import (
	_ "gfwebtest/boot"

	_ "gfwebtest/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gsession"
	"time"
)

//type ControllerTemplate struct {
//	gmvc.Controller
//}
//func (c *ControllerTemplate) Info() {
//	c.View.Assign("name", "john")
//	c.View.Assigns(map[string]interface{}{
//		"age"   : 18,
//		"score" : 100,
//	})
//	c.View.DisplayContent(`
//       <html>
//           <head>
//               <title>gf template engine</title>
//           </head>
//           <body>
//               <p>Name: {{.name}}</p>
//               <p>Age:  {{.age}}</p>
//               <p>Score:{{.score}}</p>
//           </body>
//       </html>
//   `)
//}
//
//func MiddlewareAuth(r *ghttp.Request) {
//	token := r.Get("token")
//	r.Response.Write("auth....")
//	if token == "123456" {
//		r.Middleware.Next()
//	} else {
//		r.Response.WriteStatus(http.StatusForbidden)
//	}
//}
//
//func MiddlewareCORS(r *ghttp.Request) {
//	r.Response.CORSDefault()
//	r.Middleware.Next()
//}
//
//func MiddlewareLog(r *ghttp.Request) {
//	r.Middleware.Next()
//	g.Log().Println(r.Response.Status, r.URL.Path)
//}
//
//
//type Controller struct{
//	gmeta.Meta `group:"users" obj:"user"`
//}
//
//func (c *Controller) Index(r *ghttp.Request) {
//	r.Response.Write("index")
//}
//
//func (c *Controller) Show(r *ghttp.Request) {
//	r.Response.Write("show")
//}
func main() {
	s := g.Server()
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,
		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
	})
	s.BindHandler("/main1", func(r *ghttp.Request) {
		//panic(errors.New("test-exp-handler-error"))
		r.Response.WriteTpl("inclayout.html", g.Map{
			"header":  "incheader",
			"mainTpl": "maindir/main1.html",
			"footer":  "incfooter",
		})

	})
	s.BindHandler("/main2", func(r *ghttp.Request) {
		g.Log("info").Info("debug.... log")
		r.Response.WriteTpl("inclayout.html", g.Map{
			"header":  "inclheader2",
			"mainTpl": "maindir/main2.html",
			"footer":  "inclfooter2",
		})
	})

	//custom status hint info
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		//403 : func(r *ghttp.Request){r.Response.Writeln("403")},
		404: func(r *ghttp.Request) { r.Response.Writeln("404") },
		//500 : func(r *ghttp.Request){r.Response.Writeln("500")},
	})
	s.SetIndexFolder(true)
	s.SetServerRoot("/home/jy/")
	s.Run()
	/*s.BindHandler("/cs", func(r *ghttp.Request) {
		r.Cookie.Set("theme", "default")
		r.Session.Set("name", "john")
		r.Response.WriteTplContent(`Cookie:{{.Cookie.theme}}, Session:{{.Session.name}}`, nil)
	})
	s.BindHandler("/",func(r *ghttp.Request){
		r.Response.WriteTpl("layout.html", g.Map{
			"header":    "this is header",
			"container": "This is container",
			"footer":    "This is footer",
		})
	})
	s.BindHandler("/main1",func(r *ghttp.Request){
		r.Response.WriteTpl("inclayout.html", g.Map{
			"header":    "incheader",
			"mainTpl":   "maindir/main1.html",
			"footer":    "incfooter",
		})
	})
	s.BindHandler("/main2",func(r *ghttp.Request){
		r.Response.WriteTpl("inclayout.html", g.Map{
			"header":      "inclheader2",
			"mainTpl":    "maindir/main2.html",
			"footer":       "inclfooter2",
		})
	})
	s.SetIndexFolder(true)
	s.SetServerRoot("/home/jy/soft")

	s.BindHandler("/{class}-{course}/:name/*act", func(r *ghttp.Request) {
		r.Response.Writef(
			"%v %v %v %v",
			r.Get("class"),
			r.Get("course"),
			r.Get("name"),
			r.Get("act"),
		)
	})

	//s.BindHandler("/:name", func(r *ghttp.Request){
	//	r.Response.Writeln(r.Router.Uri)
	//})
	s.BindHandler("/:name/update", func(r *ghttp.Request){
		r.Response.Writeln(r.Router.Uri)
	})
	//s.BindHandler("/:name/:action", func(r *ghttp.Request){
	//	r.Response.Writeln(r.Router.Uri)
	////})
	//s.BindHandler("/:name/*any", func(r *ghttp.Request){
	//	r.Response.Writeln(r.Router.Uri)
	//})
	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request){
		r.Response.Writeln(r.Router.Uri)
	})

	s.BindObject("/api",new (Controller))

	s.Use(MiddlewareLog)
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareAuth, MiddlewareCORS)
		group.GET("/test", func(r *ghttp.Request) {
			r.Response.Write("test")
		})
		group.Group("/order", func(group *ghttp.RouterGroup) {
			group.GET("/list", func(r *ghttp.Request) {
				r.Response.Write("list")
			})
			group.PUT("/update", func(r *ghttp.Request) {
				r.Response.Write("update")
			})
		})
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.GET("/info", func(r *ghttp.Request) {
				r.Response.Write("info")
			})
			group.POST("/edit", func(r *ghttp.Request) {
				r.Response.Write("edit")
			})
			group.DELETE("/drop", func(r *ghttp.Request) {
				r.Response.Write("drop")
			})
		})
		group.Group("/hook", func(group *ghttp.RouterGroup) {
			group.Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.Response.Write("hook any\n")
			})
			group.Hook("/:name", ghttp.HookBeforeServe, func(r *ghttp.Request) {
				r.Response.Write("hook name\n")
			})
		})
	})*/

}
