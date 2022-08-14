package base

import "context"

type AppServiceInterface interface {
	FetchApp(context.Context, interface{}) interface{}
}
