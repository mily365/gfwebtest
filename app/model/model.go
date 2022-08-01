// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// ControlInfo is the golang structure for table control_info.
type ControlInfo struct {
	Id              int         `orm:"id,primary"       json:"id"`              //
	Title           string      `orm:"title"            json:"title"`           //
	PropName        string      `orm:"prop_name"        json:"propName"`        //
	Icon            string      `orm:"icon"             json:"icon"`            //
	Order           string      `orm:"order"            json:"order"`           //
	GroupName       string      `orm:"group_name"       json:"groupName"`       //
	IsKey           int         `orm:"is_key"           json:"isKey"`           //
	IsHidden        int         `orm:"is_hidden"        json:"isHidden"`        //
	ControlType     string      `orm:"control_type"     json:"controlType"`     //
	ValidatorType   string      `orm:"validator_type"   json:"validatorType"`   //
	ControlPosition string      `orm:"control_position" json:"controlPosition"` //
	ExtraInfo       string      `orm:"extra_Info"       json:"extraInfo"`       //
	CreatedTime     *gtime.Time `orm:"created_time"     json:"createdTime"`     //
	UpdatedTime     *gtime.Time `orm:"updated_time"     json:"updatedTime"`     //
	DeletedTime     *gtime.Time `orm:"deleted_time"     json:"deletedTime"`     //
	Version         int         `orm:"version"          json:"version"`         //
	Sid             int         `orm:"sid"              json:"sid"`             //
	IsSort          int         `orm:"is_sort"          json:"isSort"`          //
	SqlType         string      `orm:"sql_type"         json:"sqlType"`         //
	SqlLength       int         `orm:"sql_length"       json:"sqlLength"`       //
	SqlDefault      string      `orm:"sql_default"      json:"sqlDefault"`      //
	SqlName         string      `orm:"sql_name"         json:"sqlName"`         //
	IsQuickSearch   string      `orm:"is_quick_search"  json:"isQuickSearch"`   //
}

// Project is the golang structure for table project.
type Project struct {
	Id   int    `orm:"id,primary" json:"id"`   //
	Name string `orm:"name"       json:"name"` //
	Age  int    `orm:"age"        json:"age"`  // 年龄
}

// Solution is the golang structure for table solution.
type Solution struct {
	Id          int         `orm:"id,primary"   json:"id"`          //
	BizCode     string      `orm:"biz_code"     json:"bizCode"`     // 业务编码，资源编码
	ModelName   string      `orm:"model_name"   json:"modelName"`   // 方案所描述的实体名称，后端实体名称
	Title       string      `orm:"title"        json:"title"`       // 方案显示名称
	CreatedTime *gtime.Time `orm:"created_time" json:"createdTime"` //
	UpdatedTime *gtime.Time `orm:"updated_time" json:"updatedTime"` //
	DeletedTime *gtime.Time `orm:"deleted_time" json:"deletedTime"` //
	Version     int         `orm:"version"      json:"version"`     //
	Lang        string      `orm:"lang"         json:"lang"`        // 语言编码
}

// User is the golang structure for table user.
type User struct {
	Id        uint        `orm:"id,primary" json:"id"`        //
	Name      string      `orm:"name"       json:"name"`      //
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
	Version   int         `orm:"version"    json:"version"`   //
	Age       int         `orm:"age"        json:"age"`       //
}

// UserDetail is the golang structure for table user_detail.
type UserDetail struct {
	Uid     uint   `orm:"uid,primary" json:"uid"`     //
	Address string `orm:"address"     json:"address"` //
}

// UserScore is the golang structure for table user_score.
type UserScore struct {
	Id     uint   `orm:"id,primary" json:"id"`     //
	Uid    uint   `orm:"uid"        json:"uid"`    //
	Score  uint   `orm:"score"      json:"score"`  //
	Course string `orm:"course"     json:"course"` //
}
