
package model
import "xpass/app"
func NewApp() interface{}{
	var app *App
	return &app
}
func NewApps() interface{}{
	var apps []*App
	return &apps
}
func init() {
	fun:=NewApp
	funs:=NewApps
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("app", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("apps", funs)
}
