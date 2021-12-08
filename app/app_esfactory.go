package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/olivere/elastic/v7"
	"reflect"
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

//按照输入模糊查询
func (esF *esFactory) CreateIndex(ctx context.Context, modelType reflect.Type) string {
	esName, appName := GetEsName(modelType.Name())
	//创建索引模板，日期
	appIndexTmplName := appName + "index__tmpl"
	isTmplExist, _ := esF.Client.IndexTemplateExists(appIndexTmplName).Do(ctx)
	if isTmplExist == false {
		mapIndexTmpl := g.Map{
			"index_patterns": []string{appName + "*"},
			"template": g.Map{
				"mappings": g.Map{
					"dynamic_date_formats": []string{"yyyy-MM-dd HH:mm:ss", "yyyy-MM-dd HH:mm:ss.SSS"},
				},
			},
		}
		_, errTmp := esF.Client.IndexPutIndexTemplate(appIndexTmplName).BodyJson(mapIndexTmpl).Do(ctx)
		if errTmp != nil {
			panic(errTmp.Error())
		}
		//_, err2 := esF.Client.Flush().Do(ctx)
		//if err2 != nil {
		//	panic(err2)
		//}
	}
	isExist, _ := esF.Client.IndexExists(esName).Do(ctx)
	if isExist == false {
		_, cErr := esF.Client.CreateIndex(esName).Do(ctx)
		if cErr != nil {
			panic(cErr.Error())
		}
		// Flush to make sure the documents got written.
		_, err2 := esF.Client.Flush().Index(esName).Do(ctx)
		if err2 != nil {
			panic(err2)
		}
	}
	return esName
}

func (esF *esFactory) Create(ctx context.Context, id string, jsonStr string, modelName string) interface{} {
	appNamePrefix := g.Config().GetString("appInfo.name")
	indexName := fmt.Sprintf("%s_%s", appNamePrefix, modelName)
	res, err := esF.Client.Index().Index(indexName).Id(id).BodyString(jsonStr).Do(ctx)
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

func (esF *esFactory) All(ctx context.Context, condition interface{}, modelName string) *elastic.SearchResult {
	var resultRes *elastic.SearchResult
	search := condition.(g.Map)
	skip, pageSize := PageParam(search)
	esName, _ := GetEsName(modelName)
	if search["queryForm"] != nil && (StrKeyMapIsEmpty(search["queryForm"].(g.Map)) == false) {
		query := elastic.NewBoolQuery()
		for k, v := range search["queryForm"].(g.Map) {
			query.Must(elastic.NewTermQuery(k, v))
		}
		resultRes, _ = esF.Client.Search().Index(esName).From(skip).Size(pageSize).Query(query).Do(ctx)
	} else {
		Logger.Debug(esName, "search form is empty")
		resultRes, _ = esF.Client.Search().Index(esName).From(skip).Size(pageSize).Do(ctx)
	}
	return resultRes
}

func (esF *esFactory) ScrollPage(ctx context.Context, condition interface{}, modelName string) *elastic.SearchResult {
	var resultRes *elastic.SearchResult
	appNamePrefix := g.Config().GetString("appInfo.name")
	indexName := fmt.Sprintf("%s_%s", appNamePrefix, modelName)

	search := condition.(g.Map)
	//_, pageSize := PageParam(search)
	scSV := esF.Client.Scroll(indexName).KeepAlive("5m").Size(2)
	if search["scrollId"] != nil {
		scSV = scSV.ScrollId(search["scrollId"].(string))
	}

	if search["queryForm"] != nil && (StrKeyMapIsEmpty(search["queryForm"].(g.Map)) == false) {
		Logger.Debug(indexName, "search form is not empty")
		query := elastic.NewBoolQuery()
		for k, v := range search["queryForm"].(g.Map) {
			query.Must(elastic.NewTermQuery(k, v))
		}
		resultRes, _ = scSV.Query(query).Do(ctx)
	} else {
		Logger.Debug(indexName, "search form is empty")
		resultRes, _ = scSV.Do(ctx)
	}
	Logger.Debug("scollid is ", resultRes.ScrollId)
	g.Dump(resultRes.Hits.Hits)
	return resultRes
}

func StrKeyMapIsEmpty(m g.Map) bool {
	var rtn bool = true
	for k, _ := range m {
		if k != "" {
			rtn = false
			break
		}

	}
	return rtn

}
func TranEsResultToRows(resultRes *elastic.SearchResult, modelKey string) []interface{} {
	var rows []interface{}
	for _, hit := range resultRes.Hits.Hits {
		sp := TypePointerFuncFactory.GetStructPointer(modelKey)
		err := json.Unmarshal(hit.Source, &sp)
		if err != nil {
			panic(err.Error())
		}
		rows = append(rows, sp)
	}
	return rows
}
