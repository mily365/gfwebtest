package model

import "xpass/app"

func NewSolution() interface{} {
	var solution *Solution
	return &solution
}
func NewSolutions() interface{} {
	var solutions []*Solution
	return &solutions
}

//cc
func init() {
	fun := NewSolution
	funs := NewSolutions
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("solution", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("solutions", funs)
}
