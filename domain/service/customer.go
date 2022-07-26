package service

import (
	"context"
	"hexagonal_project/adapter/driven/repository"
	"hexagonal_project/domain/entity"
	"hexagonal_project/infrastructure/model"
	"hexagonal_project/port/repository_port"
	"hexagonal_project/port/service_port"
	"sync"
	"time"
)

var (
	customerServiceOnce sync.Once
	customerServiceImpl service_port.CustomerServicePort
)

type customerService struct {
	customerRepo repository_port.CustomerRepositoryPort  // 依赖倒置
	customerAddressRepo repository_port.CustomerAddressRepositoryPort  // 依赖倒置
}

var _ service_port.CustomerServicePort = &customerService{}

func NewCustomerService() service_port.CustomerServicePort {
	customerServiceOnce.Do(func() {
		customerServiceImpl = &customerService{
			customerRepo: repository.NewCustomerRepo(),
			customerAddressRepo: repository.NewCustomerAddressRepo(),
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
	res.CreatedAt = time.Unix(int64(resRepo.CreatedAt), 0).UTC()
	res.UpdatedAt = time.Unix(int64(resRepo.UpdatedAt), 0).UTC()
	res.CreatedBy = resRepo.CreatedBy
	res.UpdatedBy = resRepo.UpdatedBy

	res.CustomerAddresses = make([]entity.CustomerAddress, 0)

	for _, v := range resRepo.CustomerAddresses {
		customerAddress := entity.CustomerAddress{}

		customerAddress.ID = v.Id
		customerAddress.Address = v.Address
		customerAddress.Contact = v.Contact
		customerAddress.Phone = v.Phone
		customerAddress.Email = v.Email
		customerAddress.DefaultSt = v.DefaultSt
		customerAddress.CreatedAt = time.Unix(int64(v.CreatedAt), 0).UTC()
		customerAddress.UpdatedAt = time.Unix(int64(v.UpdatedAt), 0).UTC()
		customerAddress.CreatedBy = v.CreatedBy
		customerAddress.UpdatedBy = v.UpdatedBy

		res.CustomerAddresses = append(res.CustomerAddresses, customerAddress)
	}

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

		item.CreatedAt = time.Unix(int64(resInfo.CreatedAt), 0).UTC()
		item.UpdatedAt = time.Unix(int64(resInfo.UpdatedAt), 0).UTC()
		item.CreatedBy = resInfo.CreatedBy
		item.UpdatedBy = resInfo.UpdatedBy
		res = append(res, item)
	}
	return
}

func (srv *customerService) Create(ctx context.Context, data entity.Customer) (res *entity.Customer, err error) {
	// 请求处理
	item := model.Customer{}
	item.Name = data.Name
	item.CustomerAddresses = make([]model.CustomerAddress, 0)
	for _, v := range data.CustomerAddresses {
		customerAddress := model.CustomerAddress{
			Address:    v.Address,
			Contact:    v.Contact,
			Phone:      v.Phone,
			Email:      v.Email,
			DefaultSt:  v.DefaultSt,
		}
		item.CustomerAddresses = append(item.CustomerAddresses, customerAddress)
	}

	// 逻辑处理
	resRepo, err := srv.customerRepo.Create(ctx, item)

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
