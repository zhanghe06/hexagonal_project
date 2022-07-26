package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SalesOrder 订单（采购无需关注货架）
type SalesOrder struct {
	BaseModel
	Code            string           `gorm:"index:udx_code,unique;column:code;type:varchar(36);default:'';not null;comment:'订单编号'" json:"code"` // 编号
	CustomerId      uint64           `gorm:"column:customer_id;type:bigint(21) unsigned;default:0;not null;comment:'客户ID'" json:"customer_id"`  // 客户ID
	Customer        Customer         `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	AddressId       uint64           `gorm:"column:address_id;type:bigint(21) unsigned;default:0;not null;comment:'地址ID'" json:"address_id"` // 地址ID
	Address         CustomerAddress  `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	TotalPrice      uint32           `gorm:"column:total_price;type:int(11) unsigned;default:0;not null;comment:'订单总价'" json:"total_price"` // 订单总价
	OrderSt         uint8            `gorm:"column:order_st;type:tinyint(4);default:0;not null;comment:'订单状态'" json:"order_st"`             // 订单状态（0:已停用,1:已启用）
	SalesOrderItems []SalesOrderItem `gorm:"foreign_key:sales_order_id"`                                                                    // Has Many
}

func (m *SalesOrder) TableName() string {
	return "sales_order"
}

func (m *SalesOrder) BeforeCreate(tx *gorm.DB) (err error) {
	// 编码 V4 基于随机数
	m.Code = uuid.New().String()
	// 总价
	for _, v := range m.SalesOrderItems {
		m.TotalPrice += v.Quantity * v.UnitPrice
	}
	return
}

// SalesOrderItem 订单明细
type SalesOrderItem struct {
	BaseModel
	SalesOrderId uint64     `gorm:"column:order_id;type:bigint(21) unsigned;default:0;not null;comment:'订单ID'" json:"order_id"` // 订单ID
	SalesOrder   SalesOrder `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	ProductId    uint64     `gorm:"column:product_id;type:bigint(21) unsigned;default:0;not null;comment:'产品ID'" json:"product_id"` // 产品ID
	Product      Product    `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Quantity     uint32     `gorm:"column:quantity;type:int(11) unsigned;default:0;not null;comment:'数量'" json:"quantity"`     // 数量
	UnitPrice    uint32     `gorm:"column:unit_price;type:int(11) unsigned;default:0;not null;comment:'单价'" json:"unit_price"` // 单价
}

func (m *SalesOrderItem) TableName() string {
	return "sales_order_item"
}
