package codetmpl

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
	"xpass/app/api/base"
	"xpass/app/model"
)

type initDataApi struct {
	gmeta.Meta `path:"api.initdata"`
	base.ApiBase
}

var (
	InitDataApi *initDataApi
)

func init() {
	InitDataApi = &initDataApi{gmeta.Meta{}, base.ApiBase{}}
	app.AppContext.RegisterObj(InitDataApi)
}

func (cg *initDataApi) InitData(r *ghttp.Request) {
	g.Log().Debug("xxxxxxxxxxxxxxxxxxxxxxxx")
	err, _ := app.ModelFactory.TxModelActions("user", func(tx *gdb.TX, xmodel *gdb.Model) (error, interface{}) {
		for i := 1; i <= 5; i++ {
			// User.
			user := model.User{
				Name: fmt.Sprintf(`name_%d`, i),
			}
			lastInsertId, err := xmodel.TX(tx).Data(user).OmitEmpty().InsertAndGetId()
			if err != nil {
				return err, nil
			}
			// Detail.
			userDetail := model.UserDetail{
				Uid:     uint(lastInsertId),
				Address: fmt.Sprintf(`address_%d`, lastInsertId),
			}
			_, err = app.ModelFactory.GetModel("user_detail").TX(tx).Data(userDetail).OmitEmpty().Insert()
			if err != nil {
				return err, nil
			}
			// Scores.
			for j := 1; j <= 5; j++ {
				userScore := model.UserScore{
					Uid:    uint(lastInsertId),
					Score:  uint(j),
					Course: string(j),
				}
				_, err = app.ModelFactory.GetModel("user_score").TX(tx).Data(userScore).OmitEmpty().Insert()
				if err != nil {
					return err, nil
				}
			}
		}
		return nil, nil
	})
	if err != nil {
		panic(err.Error())
	}
	app.WrapSuccessRtn("init data ok", "ok", r)
}
