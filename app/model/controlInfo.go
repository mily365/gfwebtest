
package model
import "xpass/app"
func NewControlInfo() interface{}{
	var controlInfo *ControlInfo
	return &controlInfo
}
func NewControlInfos() interface{}{
	var controlInfos []*ControlInfo
	return &controlInfos
}
func init() {
	fun:=NewControlInfo
	funs:=NewControlInfos
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("controlInfo", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("controlInfos", funs)
}
