
package model

import (
	"gfwebtest/app"
	"github.com/gogf/gf/util/gmeta"
)
func NewUser() interface{}{
	var user *User
	return &user
}
func NewUsers() interface{}{
	var users []*User
	return &users
}
func init() {
	fun:=NewUser
	funs:=NewUsers
	allsfun:=NewUserWithRelations
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("user", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("users", funs)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userWithDetailAndScores", allsfun)
}
func NewUserWithRelations() interface{}{
	var users []*UserWithDetailAndScore
	return &users
}

type UserWithDetailAndScore struct {
	gmeta.Meta `a:"User_id#" b:"UserDetail_uid#User_Id" c:"UserScore_uid#User_Id"`
	User *User
	UserDetail *UserDetail
	UserScore []*UserScore
}



