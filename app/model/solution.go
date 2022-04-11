package model

import (
	"github.com/gogf/gf/util/gmeta"
	"xpass/app"
)

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
	allsFun := NewSolutionWithRelations
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("solution", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("solutions", funs)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("solutionWithControlInfos", allsFun)
}
func NewSolutionWithRelations() interface{} {
	var sls []*SolutionWithControlInfo
	return &sls
}

type SolutionWithControlInfo struct {
	gmeta.Meta  `a:"Solution_Id#" b:"ControlInfo_Sid#Solution_Id"`
	Solution    *Solution
	ControlInfo []*ControlInfo
}
