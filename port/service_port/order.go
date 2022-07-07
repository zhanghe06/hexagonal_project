package service_port

import (
	"context"
	"hexagonal_project/domain/entity"
)

// OrderServicePort 驱动端口 内部实现
type OrderServicePort interface {
	GetInfo(ctx context.Context, id uint64) (res *entity.Order, err error)
	//Create(ctx context.Context, data entity.Order) (res *entity.Order, err error)
	//Update(ctx context.Context, id uint64, data map[string]interface{}) (res *entity.Order, err error)
	//GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []entity.Order, err error)
}
