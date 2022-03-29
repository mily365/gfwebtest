package model

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
)

func NewUser() interface{} {
	var user *User
	return &user
}
func NewUsers() interface{} {
	var users []*User
	return &users
}
func init() {
	fun := NewUser
	funs := NewUsers
	allsfun := NewUserWithRelations
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("user", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("users", funs)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userWithDetailAndScores", allsfun)
}
func NewUserWithRelations() interface{} {
	var users []*UserWithDetailAndScore
	return &users
}

type UserWithDetailAndScore struct {
	gmeta.Meta `a:"User_Id#" b:"UserDetail_Uid#User_Id" c:"UserScore_Uid#User_Id"`
	User       *User
	UserDetail *UserDetail
	UserScore  []*UserScore
}
