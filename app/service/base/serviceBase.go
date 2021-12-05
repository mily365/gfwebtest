package base

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"xpass/app"
)

type ServiceBase struct {
	Dao interface{}
}

func (s *ServiceBase) Withalls(ctx context.Context, i interface{}) interface{} {
	return s.Dao.(app.CommonOperation).Withalls(ctx, i)
}
func (s *ServiceBase) Scrollpage(ctx context.Context, i interface{}) interface{} {
	g.Log().Debug("service Scrollpage")
	return s.Dao.(app.CommonOperation).Scrollpage(ctx, i)
}
func (s *ServiceBase) All(ctx context.Context, i interface{}) interface{} {
	return s.Dao.(app.CommonOperation).All(ctx, i)
}

func (s *ServiceBase) Create(ctx context.Context, i interface{}) interface{} {
	return s.Dao.(app.CommonOperation).Create(ctx, i)
}

func (s *ServiceBase) Update(ctx context.Context, i interface{}) interface{} {
	return s.Dao.(app.CommonOperation).Update(ctx, i)
}

func (s *ServiceBase) Delete(ctx context.Context, i interface{}) interface{} {
	return s.Dao.(app.CommonOperation).Delete(ctx, i)
}
