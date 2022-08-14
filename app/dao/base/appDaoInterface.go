package base

import (
	"context"
)

type AppDaoInterface interface {
	FetchApp(context.Context, interface{}) interface{}
}
