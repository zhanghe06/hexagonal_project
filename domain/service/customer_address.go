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
	customerAddressServiceOnce sync.Once
	customerAddressServiceImpl service_port.CustomerAddressServicePort
)

type customerAddressService struct {
	customerAddressRepo repository_port.CustomerAddressRepositoryPort  // 依赖倒置
}

var _ service_port.CustomerAddressServicePort = &customerAddressService{}

func NewCustomerAddressService() service_port.CustomerAddressServicePort {
	customerAddressServiceOnce.Do(func() {
		customerAddressServiceImpl = &customerAddressService{
			customerAddressRepo: repository.NewCustomerAddressRepo(),
		}
	})
	return customerAddressServiceImpl
}

func (srv *customerAddressService) GetInfo(ctx context.Context, id uint64) (res *entity.CustomerAddress, err error) {
	// 逻辑处理
	resRepo, err := srv.customerAddressRepo.GetInfo(ctx, id)

	if err != nil {
		return
	}

	// 响应处理
	res = &entity.CustomerAddress{}
	res.ID = resRepo.Id
	res.Address = resRepo.Address
	res.Contact = resRepo.Contact
	res.Phone = resRepo.Phone
	res.Email = resRepo.Email
	res.DefaultSt = resRepo.DefaultSt

	res.CreatedAt = time.Unix(int64(resRepo.CreatedAt), 0).UTC()
	res.UpdatedAt = time.Unix(int64(resRepo.UpdatedAt), 0).UTC()
	res.CreatedBy = resRepo.CreatedBy
	res.UpdatedBy = resRepo.UpdatedBy

	return
}

func (srv *customerAddressService) GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []entity.CustomerAddress, err error) {
	total, resList, err := srv.customerAddressRepo.GetList(ctx, filter, args...)
	// 响应处理
	res = make([]entity.CustomerAddress, 0)
	for _, resInfo := range resList {
		item := entity.CustomerAddress{}
		item.ID = resInfo.Id
		item.Address = resInfo.Address
		item.Contact = resInfo.Contact
		item.Phone = resInfo.Phone
		item.Email = resInfo.Email
		item.DefaultSt = resInfo.DefaultSt

		item.CreatedAt = time.Unix(int64(resInfo.CreatedAt), 0).UTC()
		item.UpdatedAt = time.Unix(int64(resInfo.UpdatedAt), 0).UTC()
		item.CreatedBy = resInfo.CreatedBy
		item.UpdatedBy = resInfo.UpdatedBy

		res = append(res, item)
	}
	return
}

func (srv *customerAddressService) Create(ctx context.Context, data entity.CustomerAddress) (res *entity.CustomerAddress, err error) {
	// 请求处理
	item := model.CustomerAddress{}
	item.Address = data.Address
	item.Contact = data.Contact
	item.Phone = data.Phone
	item.Email = data.Email
	item.DefaultSt = data.DefaultSt

	// 逻辑处理
	resRepo, err := srv.customerAddressRepo.Create(ctx, item)

	// 响应处理
	if err != nil {
		return
	}
	res = &data
	res.ID = resRepo.Id

	return
}

func (srv *customerAddressService) Update(ctx context.Context, id uint64, data map[string]interface{}) (err error) {
	// 逻辑处理
	err = srv.customerAddressRepo.Update(ctx, id, data)
	return
}

func (srv *customerAddressService) Delete(ctx context.Context, id uint64) (err error) {
	// 逻辑处理
	err = srv.customerAddressRepo.Delete(ctx, id)
	return
}
