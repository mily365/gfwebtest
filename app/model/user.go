
package model
import "xpass/app"
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
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("user", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("users", funs)
}
