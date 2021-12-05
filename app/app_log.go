package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
)

// Logger ----------------------------------
var Logger = g.Log()

func LoggerWithCtx(ctx context.Context) *glog.Logger {
	Logger.SetHandlers(LoggingJsonHandler)
	return Logger.Ctx(ctx)
}

// JsonOutputsForLogger LoggingJsonHandler is a example handler for logging JSON format content.
type JsonOutputsForLogger struct {
	TraceId string `json:"traceId"`
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}

var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	if ctx.Value(TraceID) != nil {
		jsonForLogger := JsonOutputsForLogger{
			TraceId: ctx.Value(TraceID).(string),
			Time:    in.TimeFormat,
			Level:   gstr.Trim(in.LevelFormat, "[]"),
			Content: gstr.Trim(in.Content),
		}
		jsonBytes, err := json.Marshal(jsonForLogger)
		if err != nil {
			panic(errors.New("json log handler error!"))
		}
		if in.Level == glog.LEVEL_ERRO {
			// to es error
			fmt.Print("error.......")
		}
		if in.Level == glog.LEVEL_INFO {
			// to es
			fmt.Print("info.......")
		}
		in.Buffer().Write(jsonBytes)
		//in.Buffer().WriteString("\n")
		in.Content = string(jsonBytes)
	}

	//fmt.Println("to push es....")
	//fmt.Println(string(jsonBytes))

	//in.Content=string(jsonBytes)+"\n"
	//to do write to es......
	in.Next()

}
