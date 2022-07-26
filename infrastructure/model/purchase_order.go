package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PurchaseOrder 订单（采购无需关注货架）
type PurchaseOrder struct {
	BaseModel
	Code               string              `gorm:"index:udx_code,unique;column:code;type:varchar(36);default:'';not null;comment:'订单编号'" json:"code"` // 订单编号
	SupplierId         uint64              `gorm:"column:supplier_id;type:bigint(21) unsigned;default:0;not null;comment:'供方ID'" json:"supplier_id"`  // 供方ID
	Supplier           Supplier            `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	AddressId          uint64              `gorm:"column:address_id;type:bigint(21) unsigned;default:0;not null;comment:'地址ID'" json:"address_id"` // 地址ID
	Address            SupplierAddress     `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	TotalPrice         uint32              `gorm:"column:total_price;type:int(11) unsigned;default:0;not null;comment:'订单总价'" json:"total_price"` // 订单总价
	OrderSt            uint8               `gorm:"column:order_st;type:tinyint(4);default:0;not null;comment:'订单状态'" json:"order_st"`             // 订单状态（0:已停用,1:已启用）
	PurchaseOrderItems []PurchaseOrderItem `gorm:"foreign_key:purchase_order_id"`                                                                 // Has Many
}

func (m *PurchaseOrder) TableName() string {
	return "purchase_order"
}

func (m *PurchaseOrder) BeforeCreate(tx *gorm.DB) (err error) {
	// 编码 V4 基于随机数
	m.Code = uuid.New().String()
	// 总价
	for _, v := range m.PurchaseOrderItems {
		m.TotalPrice += v.Quantity * v.UnitPrice
	}
	return
}

// PurchaseOrderItem 订单明细
type PurchaseOrderItem struct {
	BaseModel
	PurchaseOrderId uint64        `gorm:"column:order_id;type:bigint(21) unsigned;default:0;not null;comment:'订单ID'" json:"order_id"` // 订单ID
	PurchaseOrder   PurchaseOrder `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	ProductId       uint64        `gorm:"column:product_id;type:bigint(21) unsigned;default:0;not null;comment:'产品ID'" json:"product_id"` // 产品ID
	Product         Product       `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Quantity        uint32        `gorm:"column:quantity;type:int(11) unsigned;default:0;not null;comment:'数量'" json:"quantity"`     // 数量
	UnitPrice       uint32        `gorm:"column:unit_price;type:int(11) unsigned;default:0;not null;comment:'单价'" json:"unit_price"` // 单价
}

func (m *PurchaseOrderItem) TableName() string {
	return "purchase_order_item"
}
