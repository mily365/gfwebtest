package base

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"xpass/app"
)

type ServiceBase struct {
	Dao app.CommonOperation
}

func (s *ServiceBase) Withalls(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Withalls(ctx, i)
}
func (s *ServiceBase) Scrollpage(ctx context.Context, i interface{}) interface{} {
	g.Log().Debug("service Scrollpage")
	return s.Dao.Scrollpage(ctx, i)
}
func (s *ServiceBase) All(ctx context.Context, i interface{}) interface{} {
	return s.Dao.All(ctx, i)
}

func (s *ServiceBase) Create(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Create(ctx, i)
}

func (s *ServiceBase) Update(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Update(ctx, i)
}

func (s *ServiceBase) Delete(ctx context.Context, i interface{}) interface{} {
	return s.Dao.Delete(ctx, i)
}
