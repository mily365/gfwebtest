package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/guid"
)

// Logger ----------------------------------
var Logger = g.Log()

func init() {
	Logger.SetHandlers(LoggingJsonHandler)
}

func LoggerWithCtx(ctx context.Context) *glog.Logger {
	return Logger.Ctx(ctx)
}

// JsonOutputsForLogger LoggingJsonHandler is a example handler for logging JSON format content.
type JsonOutputsForLogger struct {
	BackendTime int64  `json:"backendTime"`
	TraceId     string `json:"traceId"`
	Time        string `json:"time"`
	Level       string `json:"level"`
	Content     string `json:"content"`
}

var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	if ctx.Value(TraceID) != nil {
		backendTime := gtime.TimestampMilli() - ctx.Value(ResponseTimeKey).(int64)
		jsonForLogger := JsonOutputsForLogger{
			BackendTime: backendTime,
			TraceId:     ctx.Value(TraceID).(string),
			Time:        gtime.Now().UTC().Format("Y-m-d H:i:s.u"),
			Level:       gstr.Trim(in.LevelFormat, "[]"),
			Content:     gstr.Trim(in.Content),
		}
		jsonBytes, err := json.Marshal(jsonForLogger)
		if err != nil {
			panic(errors.New("json log handler error!"))
		}
		if in.Level == glog.LEVEL_ERRO {
			// to es error
			fmt.Print("error.......")
			if g.Config().GetBool("app.enableEs") == true {
				GetEsFactory().Create(context.Background(), guid.S(), string(jsonBytes), "error_log")
			}

		}
		if in.Level == glog.LEVEL_INFO {
			// to es
			g.Dump("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			if g.Config().GetBool("app.enableEs") == true {
				GetEsFactory().Create(context.Background(), guid.S(), string(jsonBytes), "info_log")
			}

			fmt.Print("info.......")
		}
		in.Buffer().Write(jsonBytes)
		in.Buffer().WriteString("\n")
		//in.Content = string(jsonBytes)
	}

	//fmt.Println("to push es....")
	//fmt.Println(string(jsonBytes))

	//in.Content=string(jsonBytes)+"\n"
	//to do write to es......
	in.Next()

}
