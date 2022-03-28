package middleware

import (
	"fmt"
	"xpass/app"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"net/http"
)

var (
	Auth *auth
)

func init() {
	Auth = new(auth)
	app.AppContext.RegisterObj(Auth)
}

type auth struct {
	gmeta.Meta `path:"middle.api"`
}

// 鉴权中间件，只有登录成功之后才能通过
func (s *auth) Auth(r *ghttp.Request) {
	fmt.Println("auth called....................................")
	r.Response.WriteStatus(http.StatusForbidden)
	r.Middleware.Next()
}
