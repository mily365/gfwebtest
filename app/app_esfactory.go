package app

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/olivere/elastic/v7"
	"time"
)

//--------------es的客户端工厂----------------------------
var esClientFactory = &esFactory{}

type esFactory struct {
	Client *elastic.Client
}

func GetEsFactory() *esFactory {
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

func (esF *esFactory) Create(ctx context.Context, id string, data interface{}, modelName string) interface{} {
	appNamePrefix := g.Config().GetString("appInfo.name")
	indexName := fmt.Sprintf("%s_%s", appNamePrefix, modelName)
	res, err := esF.Client.Index().Index(indexName).Id(id).BodyJson(data).Do(ctx)
	if err != nil {
		panic(err)
	}
	// Flush to make sure the documents got written.
	_, err2 := esF.Client.Flush().Index(indexName).Do(ctx)
	if err2 != nil {
		panic(err2)
	}
	return res
}
func (esF *esFactory) ScrollPage(ctx context.Context, condition interface{}, modelName string) interface{} {
	appNamePrefix := g.Config().GetString("appInfo.name")
	indexName := fmt.Sprintf("%s_%s", appNamePrefix, modelName)
	Logger.Debug(indexName, "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	search := condition.(g.Map)
	_, pageSize := PageParam(search)
	scSV := esF.Client.Scroll(indexName).Size(pageSize)
	query := elastic.NewBoolQuery()
	for k, v := range search["queryForm"].(g.Map) {
		query.Must(elastic.NewTermQuery(k, v))
	}
	resultRes, _ := scSV.Query(query).Do(ctx)
	return resultRes
}
