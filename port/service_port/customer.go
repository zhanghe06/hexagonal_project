package service_port

import (
	"context"
	"hexagonal_project/domain/entity"
)

// CustomerServicePort 内部实现
type CustomerServicePort interface {
	GetInfo(ctx context.Context, id uint64) (res *entity.Customer, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []entity.Customer, err error)
	Create(ctx context.Context, data entity.Customer) (res *entity.Customer, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
