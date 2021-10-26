package model

import "github.com/gogf/gf/net/ghttp"

const (
	ContextInfoKey = "ContextInfoKey"
)
type ContextInfo struct {
	Session *ghttp.Session // 当前Session管理对象
	User    *ContextUser   // 上下文用户信息
}
// 请求上下文中的用户信息
type ContextUser struct {
	Id       int   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
}
