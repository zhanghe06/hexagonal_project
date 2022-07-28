package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"hexagonal_project/infrastructure/db"
	"hexagonal_project/infrastructure/model"
	"hexagonal_project/port/repository_port"
	"strings"
	"sync"
	"time"
)

var (
	customerRepoOnce sync.Once
	customerRepoImpl repository_port.CustomerRepositoryPort
)

type customerRepo struct {
	db *gorm.DB
}

var _ repository_port.CustomerRepositoryPort = &customerRepo{}

func NewCustomerRepo() repository_port.CustomerRepositoryPort {
	customerRepoOnce.Do(func() {
		customerRepoImpl = &customerRepo{
			db: db.NewDB(),
		}
	})
	return customerRepoImpl
}

func (repo *customerRepo) GetInfo(ctx context.Context, id uint64) (res *model.Customer, err error) {
	tx := repo.db.WithContext(ctx)

	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id

	err = tx.Where(condition).First(&res).Error
	if err != nil {
		return
	}

	err = tx.Model(&res).Association("CustomerAddresses").Find(&res.CustomerAddresses)
	if err != nil {
		return
	}

	return
}

func (repo *customerRepo) GetList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []model.Customer, err error) {
	tx := repo.db.WithContext(ctx)

	// 排序条件
	sorter := fmt.Sprintf("%s %s", "id", "DESC")
	if v, ok := filter["sorter"]; ok {
		sortOrder := strings.Split(v.(string), " ")
		if sortOrder[1] == "ascend" {
			sorter = fmt.Sprintf("%s %s", sortOrder[0], "ASC")
		} else if sortOrder[1] == "descend" {
			sorter = fmt.Sprintf("%s %s", sortOrder[0], "DESC")
		} else {
			sorter = v.(string)
		}
		delete(filter, "sorter")
	}

	// 查询条件
	limit := 10
	offset := 0
	condition := make(map[string]interface{})
	for k, v := range filter {
		// 分页条件
		if k == "limit" {
			limit = int(v.(float64))
		} else if k == "offset" {
			offset = int(v.(float64))
		} else {
			condition[k] = v
		}
	}

	dbQuery := tx.Model(&model.Customer{}).Where(condition)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	// 执行排序
	dbQuery = dbQuery.Order(sorter)

	// 总记录数
	dbQuery.Count(&total)

	// 分页查询
	err = dbQuery.Limit(limit).Offset(offset).Find(&res).Error
	return
}

func (repo *customerRepo) Create(ctx context.Context, data model.Customer) (res *model.Customer, err error) {
	tx := repo.db.WithContext(ctx)

	// 逻辑处理
	result := tx.Create(&data)

	// 响应处理
	err = result.Error
	if err == nil {
		res = &data
	}
	return
}

func (repo *customerRepo) Update(ctx context.Context, id uint64, data map[string]interface{}) (err error) {
	tx := repo.db.WithContext(ctx)

	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id

	var customer model.Customer
	err = tx.Where(condition).First(&customer).Error
	if err != nil {
		return
	}

	err = tx.Model(&customer).Updates(data).Error
	return
}

func (repo *customerRepo) Delete(ctx context.Context, id uint64) (err error) {
	tx := repo.db.WithContext(ctx)

	// 参数处理
	userId := ctx.Value("userId")
	if userId == nil {
		userId = 0
	}
	operator := uint64(userId.(int))
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id

	var customer model.Customer
	err = tx.Where(condition).First(&customer).Error
	if err != nil {
		return
	}

	// 逻辑删除
	customerUpdate := model.Customer{}
	customerUpdate.DeletedAt = soft_delete.DeletedAt(time.Now().Unix())
	customerUpdate.DeletedBy = operator

	// Without Hooks/Time Tracking: https://gorm.io/docs/update.html#Without-Hooks-x2F-Time-Tracking
	err = tx.Model(&customer).UpdateColumns(customerUpdate).Error

	return
}
