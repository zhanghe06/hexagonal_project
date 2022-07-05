package model

// Order 订单
type Order struct {
	BaseModel
	CustomerId uint64 `gorm:"column:customer_id;type:bigint(21) unsigned;NOT NULL" json:"customer_id"` // 客户ID
	AddressId  uint64 `gorm:"column:address_id;type:bigint(21) unsigned;NOT NULL" json:"address_id"`   // 地址ID
	TotalPrice string `gorm:"column:total_price;type:varchar(32);NOT NULL" json:"total_price"`         // 总价
	OrderState uint8  `gorm:"column:order_st;type:tinyint(4);default:0;NOT NULL" json:"order_st"`      // 订单状态（0:已停用,1:已启用）
}

func (m *Order) TableName() string {
	return "order"
}

// OrderItems 订单明细
type OrderItems struct {
	BaseModel
	OrderId   uint64  `gorm:"column:order_id;type:bigint(21) unsigned;NOT NULL" json:"order_id"`     // 订单ID
	ProductId uint64  `gorm:"column:product_id;type:bigint(21) unsigned;NOT NULL" json:"product_id"` // 产品ID
	Quantity  uint16  `gorm:"column:quantity;type:smallint(6);NOT NULL" json:"quantity"`             // 数量
	UnitPrice float32 `gorm:"column:unit_price;type:decimal(10,2);NOT NULL" json:"unit_price"`       // 单价
}

func (m *OrderItems) TableName() string {
	return "order_items"
}
