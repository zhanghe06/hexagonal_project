package repository_port

import (
	"context"
	"hexagonal_project/infrastructure/model"
)

// OrderRepositoryPort 外部实现
type OrderRepositoryPort interface {
	GetInfo(ctx context.Context, id uint64) (res *model.Order, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.Order, err error)
	Create(ctx context.Context, data model.Order) (res *model.Order, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
