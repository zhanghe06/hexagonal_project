package model

import (
	"gorm.io/gorm"
)

// Customer 客户
type Customer struct {
	BaseModel
	Name              string            `gorm:"unique_index:udx_name;column:name;type:varchar(128);default:'';not null;comment:'客户名称'" json:"name"`
	CustomerAddresses []CustomerAddress `gorm:"foreign_key:customer_id"` // Has Many
}

func (m *Customer) TableName() string {
	return "customer"
}

func (m *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	// 参数处理
	userId := ctx.Value("user_id")
	if userId == nil {
		userId = 0
	}
	operator := userId.(uint64)
	m.CreatedBy = operator
	m.UpdatedBy = operator

	return
}

func (m *Customer) BeforeUpdate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	// 参数处理
	userId := ctx.Value("user_id")
	if userId == nil {
		userId = 0
	}
	operator := userId.(uint64)
	//m.UpdatedBy = operator // not work
	tx.Statement.SetColumn("updated_by", operator)

	return
}
