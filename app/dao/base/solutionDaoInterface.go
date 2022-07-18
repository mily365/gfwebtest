package base

import (
	"context"
	"github.com/gogf/gf/database/gdb"
)

type SolutionDaoInterface interface {
	CopyRelChildren(context.Context, interface{}, *gdb.TX, *gdb.Model) interface{}
}
