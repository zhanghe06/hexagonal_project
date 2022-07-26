package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InBound 入库（仓库看不到价格）
type InBound struct {
	BaseModel
	Code         string        `gorm:"index:udx_code,unique;column:code;type:varchar(36);default:'';not null;comment:'入库编号'" json:"code"` // 编号
	OrderId      uint64        `gorm:"column:order_id;type:bigint(21) unsigned;default:0;not null;comment:'订单ID'" json:"order_id"`        // 订单ID
	Order        PurchaseOrder `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	InBoundSt    uint8         `gorm:"column:in_bound_st;type:tinyint(4);default:0;not null;comment:'入库状态'" json:"in_bound_st"` // 入库状态（0:已停用,1:已启用）
	InBoundItems []InBoundItem `gorm:"foreign_key:in_bound_id"`                                                                 // Has Many
}

func (m *InBound) TableName() string {
	return "in_bound"
}

func (m *InBound) BeforeCreate(tx *gorm.DB) (err error) {
	m.Code = uuid.New().String() // V4 基于随机数
	return
}

// InBoundItem 入库明细
type InBoundItem struct {
	BaseModel
	InBoundId   uint64    `gorm:"column:in_bound_id;type:bigint(21) unsigned;default:0;not null;comment:'入库ID'" json:"in_bound_id"` // 入库ID
	InBound     InBound   `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	InventoryId uint64    `gorm:"column:inventory_id;type:bigint(21) unsigned;default:0;not null;comment:'存货ID'" json:"inventory_id"` // 存货ID
	Inventory   Inventory `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Quantity    uint32    `gorm:"column:quantity;type:int(11) unsigned;default:0;not null;comment:'入库数量'" json:"quantity"` // 入库数量
}

func (m *InBoundItem) TableName() string {
	return "in_bound_item"
}

// TODO 校验入库数量
