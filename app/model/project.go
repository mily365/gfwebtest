
package model
import "gfwebtest/app"
func NewProject() interface{}{
	var project *Project
	return &project
}
func NewProjects() interface{}{
	var projects []*Project
	return &projects
}
func init() {
	fun:=NewProject
	funs:=NewProjects
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("project", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("projects", funs)
}
