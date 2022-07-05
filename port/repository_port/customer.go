package repository_port

import (
	"context"
	"hexagonal_project/infrastructure/model"
)

// CustomerRepositoryPort 外部实现
type CustomerRepositoryPort interface {
	GetInfo(ctx context.Context, id uint64) (res *model.Customer, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.Customer, err error)
	Create(ctx context.Context, data model.Customer) (res *model.Customer, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
