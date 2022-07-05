package service

import (
	"context"
	"hexagonal_project/adapter/driven/repository"
	"hexagonal_project/domain/entity"
	"hexagonal_project/infrastructure/model"
	"hexagonal_project/port/repository_port"
	"hexagonal_project/port/service_port"
	"sync"
)

var (
	customerServiceOnce sync.Once
	customerServiceImpl service_port.CustomerServicePort
)

type customerService struct {
	customerRepo repository_port.CustomerRepositoryPort  // 依赖倒置
}

var _ service_port.CustomerServicePort = &customerService{}

func NewCustomerService() service_port.CustomerServicePort {
	customerServiceOnce.Do(func() {
		customerServiceImpl = &customerService{
			customerRepo: repository.NewCustomerRepo(),
		}
	})
	return customerServiceImpl
}

func (srv *customerService) GetInfo(ctx context.Context, id uint64) (res *entity.Customer, err error) {
	// 逻辑处理
	resRepo, err := srv.customerRepo.GetInfo(ctx, id)

	if err != nil {
		return
	}

	// 响应处理
	res = &entity.Customer{}
	res.ID = resRepo.Id
	res.Name = resRepo.Name

	res.CreatedAt = resRepo.CreatedAt
	res.UpdatedAt = resRepo.UpdatedAt
	res.CreatedBy = resRepo.CreatedBy
	res.UpdatedBy = resRepo.UpdatedBy

	return
}

func (srv *customerService) GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []entity.Customer, err error) {
	total, resList, err := srv.customerRepo.GetList(ctx, filter, args...)
	// 响应处理
	res = make([]entity.Customer, 0)
	for _, resInfo := range resList {
		item := entity.Customer{}
		item.ID = resInfo.Id
		item.Name = resInfo.Name

		item.CreatedAt = resInfo.CreatedAt
		item.CreatedBy = resInfo.CreatedBy
		item.UpdatedAt = resInfo.UpdatedAt
		item.UpdatedBy = resInfo.UpdatedBy
		res = append(res, item)
	}
	return
}

func (srv *customerService) Create(ctx context.Context, data entity.Customer) (res *entity.Customer, err error) {
	// 请求处理
	customerInfo := model.Customer{}
	customerInfo.Name = data.Name

	// 逻辑处理
	resRepo, err := srv.customerRepo.Create(ctx, customerInfo)

	// 响应处理
	if err != nil {
		return
	}
	res = &data
	res.ID = resRepo.Id

	return
}

func (srv *customerService) Update(ctx context.Context, id uint64, data map[string]interface{}) (err error) {
	// 逻辑处理
	err = srv.customerRepo.Update(ctx, id, data)
	return
}

func (srv *customerService) Delete(ctx context.Context, id uint64) (err error) {
	// 逻辑处理
	err = srv.customerRepo.Delete(ctx, id)
	return
}
