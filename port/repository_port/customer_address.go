package repository_port

import (
	"context"
	"hexagonal_project/infrastructure/model"
)

//go:generate mockgen -source=./customer_address.go -destination ./mock/mock_customer_address.go -package mock
// CustomerAddressRepositoryPort 从动端口 外部实现
type CustomerAddressRepositoryPort interface {
	GetInfo(ctx context.Context, id uint64) (res *model.CustomerAddress, err error)
	GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.CustomerAddress, err error)
	Create(ctx context.Context, data model.CustomerAddress) (res *model.CustomerAddress, err error)
	Update(ctx context.Context, id uint64, data map[string]interface{}) (err error)
	Delete(ctx context.Context, id uint64) (err error)
}
