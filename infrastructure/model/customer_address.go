package model

import (
	"gorm.io/gorm"
)

// CustomerAddress 客户地址（同一个客户支持多个地址）
type CustomerAddress struct {
	BaseModel
	CustomerId uint64   `gorm:"column:customer_id;type:bigint(21) unsigned;default:0;not null;comment:'客户ID'" json:"customer_id"` // 客户ID
	Customer   Customer `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Address    string   `gorm:"column:address;type:varchar(128);default:'';not null;comment:'地址'" json:"address"`      // 地址
	Contact    string   `gorm:"column:contact;type:varchar(32);default:'';not null;comment:'姓名'" json:"contact"`       // 姓名
	Phone      string   `gorm:"column:phone;type:varchar(128);default:'';not null;comment:'电话'" json:"phone"`          // 电话
	Email      string   `gorm:"column:email;type:varchar(128);default:'';not null;comment:'邮箱'" json:"email"`          // 邮箱
	DefaultSt  uint8    `gorm:"column:default_st;type:tinyint(4);default:0;not null;comment:'默认状态'" json:"default_st"` // 默认状态（0:常规,1:默认）
}

func (m *CustomerAddress) TableName() string {
	return "customer_address"
}

func (m *CustomerAddress) BeforeCreate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	// 参数处理
	userId := ctx.Value("userId")
	if userId == nil {
		userId = 0
	}
	operator := uint64(userId.(int))
	m.CreatedBy = operator
	m.UpdatedBy = operator

	return
}

func (m *CustomerAddress) BeforeUpdate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context

	// 参数处理
	userId := ctx.Value("userId")
	if userId == nil {
		userId = 0
	}
	operator := uint64(userId.(int))
	tx.Statement.SetColumn("updated_by", operator)

	return
}
