package repository

import (
	"context"
	"gorm.io/gorm"
	"hexagonal_project/infrastructure/db"
	"hexagonal_project/infrastructure/model"
	"hexagonal_project/port/repository_port"
	"sync"
)

var (
	orderRepoOnce sync.Once
	orderRepoImpl repository_port.OrderRepositoryPort
)

type orderRepo struct {
	db *gorm.DB
}

var _ repository_port.OrderRepositoryPort = &orderRepo{}

func NewOrderRepo() repository_port.OrderRepositoryPort {
	orderRepoOnce.Do(func() {
		orderRepoImpl = &orderRepo{
			db: db.NewDB(),
		}
	})
	return orderRepoImpl
}

func (repo *orderRepo) GetInfo(ctx context.Context, id uint64) (res *model.SalesOrder, err error) {
	return
}

func (repo *orderRepo) GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.SalesOrder, err error) {
	return
}

func (repo *orderRepo) Create(ctx context.Context, data model.SalesOrder) (res *model.SalesOrder, err error) {
	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return
	}

	if err = tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return
	}

	//TODO: other operation
	//if err = tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
	//	tx.Rollback()
	//	return
	//}

	err = tx.Commit().Error
	return
}

func (repo *orderRepo) Update(ctx context.Context, id uint64, data map[string]interface{}) (err error) {
	return
}

func (repo *orderRepo) Delete(ctx context.Context, id uint64) (err error) {
	return
}
