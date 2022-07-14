package base

import "context"

type SolutionInterface interface {
	CreateTable(context.Context, interface{}) interface{}
}
