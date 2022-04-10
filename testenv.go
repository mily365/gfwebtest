package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/util/gutil"
)

type x struct {
	M string `json:"m"`
	N string `json:"n"`
}

func main() {
	fmt.Println(genv.Get("GF_GCFG_PATH"))
	m := g.MapStrAny{"hello": "ok"}
	v := &x{"hello", "ok"}
	fmt.Println(gutil.Export(v))
	g.Dump(m)
	g.Dump(v)
	g.Client()

}
