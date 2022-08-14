
package model
import "xpass/app"
func NewEnumCatalog() interface{}{
	var enumCatalog *EnumCatalog
	return &enumCatalog
}
func NewEnumCatalogs() interface{}{
	var enumCatalogs []*EnumCatalog
	return &enumCatalogs
}
func init() {
	fun:=NewEnumCatalog
	funs:=NewEnumCatalogs
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("enumCatalog", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("enumCatalogs", funs)
}
