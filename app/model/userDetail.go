
package model
import "gfwebtest/app"
func NewUserDetail() interface{}{
	var userDetail *UserDetail
	return &userDetail
}
func NewUserDetails() interface{}{
	var userDetails []*UserDetail
	return &userDetails
}
func init() {
	fun:=NewUserDetail
	funs:=NewUserDetails
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userDetail", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userDetails", funs)
}
