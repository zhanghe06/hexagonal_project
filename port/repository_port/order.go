package repository_port

import (
	"context"
	"hexagonal_project/infrastructure/model"
)

// OrderRepositoryPort 从动端口 外部实现
type OrderRepositoryPort interface {
	GetInfo(ctx context.Context, id uint64) (res *model.SalesOrder, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.SalesOrder, err error)
	Create(ctx context.Context, data model.SalesOrder) (res *model.SalesOrder, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
