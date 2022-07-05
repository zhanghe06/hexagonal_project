package service

import (
	"context"
	"hexagonal_project/adapter/driven/repository"
	"hexagonal_project/domain/entity"
	"hexagonal_project/port/repository_port"
	"hexagonal_project/port/service_port"
	"sync"
)

var (
	orderServiceOnce sync.Once
	orderServiceImpl service_port.OrderServicePort
)

type orderService struct {
	orderRepo repository_port.OrderRepositoryPort
}

var _ service_port.OrderServicePort = &orderService{}

func NewOrderService() service_port.OrderServicePort {
	orderServiceOnce.Do(func() {
		orderServiceImpl = &orderService{
			orderRepo: repository.NewOrderRepo(),
		}
	})
	return orderServiceImpl
}

func (srv *orderService) GetInfo(ctx context.Context, id uint64) (res *entity.Order, err error) {
	return
}
