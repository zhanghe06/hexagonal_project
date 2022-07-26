package service_port

import (
	"context"
	"hexagonal_project/domain/entity"
)

//go:generate mockgen -source=./customer_address.go -destination ./mock/mock_customer_address.go -package mock
// CustomerAddressServicePort 驱动端口 内部实现
type CustomerAddressServicePort interface {
	GetInfo(ctx context.Context, id uint64) (res *entity.CustomerAddress, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []entity.CustomerAddress, err error)
	Create(ctx context.Context, data entity.CustomerAddress) (res *entity.CustomerAddress, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
