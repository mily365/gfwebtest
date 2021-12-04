package api

import (
	"xpass/app"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app/api/base"
)

type uploadApi struct {
	gmeta.Meta `path:"api.upload"`
	base.ApiBase
}

var (
	Upload *uploadApi
)

func init() {
	Upload = &uploadApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(Upload)
}

func (*uploadApi) UploadFile(r *ghttp.Request) {
	files := r.GetUploadFiles("upload-file")
	names, err := files.Save("/tmp/jy")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}

func (*uploadApi) UploadShow(r *ghttp.Request) {
	r.Response.WriteTpl("upload/singleupload.html")
}

// UploadShowBatch shows uploading multiple files page.
func (*uploadApi) UploadShowBatch(r *ghttp.Request) {
	r.Response.WriteTpl("upload/multiupload.html")
}
