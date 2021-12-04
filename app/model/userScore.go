package model

import "xpass/app"

func NewUserScore() interface{} {
	var userScore *UserScore
	return &userScore
}
func NewUserScores() interface{} {
	var userScores []*UserScore
	return &userScores
}
func init() {
	fun := NewUserScore
	funs := NewUserScores
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userScore", fun)
	app.TypePointerFuncFactory.RegisterOrGetTypePointer("userScores", funs)
}
