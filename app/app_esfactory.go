package app

import (
	"github.com/gogf/gf/frame/g"
	"github.com/olivere/elastic/v7"
	"time"
)

//--------------es的客户端工厂----------------------------

var esClientFactory = esFactory{}

type esFactory struct {
	Client *elastic.Client
}

func GetEsFactory() esFactory {
	if esClientFactory.Client != nil {
		return esClientFactory
	}

	urls := g.Cfg().GetStrings("esConfig.urls")
	isSniff := g.Cfg().GetBool("esConfig.isSniff")
	//checkInterval := g.Cfg().GetInt8("esConfig.healthCheckInterval")
	uname := g.Cfg().GetString("esConfig.userName")
	pwd := g.Cfg().GetString("esConfig.pwd")
	clt, err := elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetSniff(isSniff),
		elastic.SetHealthcheckInterval(30*time.Second),
		//elastic.SetMaxRetries(5),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetBasicAuth(uname, pwd))
	if err != nil {
		panic(err.Error())
	} else {
		esClientFactory.Client = clt
	}
	return esClientFactory
}
