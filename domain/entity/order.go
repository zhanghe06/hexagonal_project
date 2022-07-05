package entity

type Order struct {
	Entity
	ID         uint64      `json:"id"`
	Customer   Customer    `json:"customer"`
	Address    Address     `json:"address"`
	TotalPrice float32     `json:"total_price"`
	OrderState uint8       `json:"order_state"`
	OrderItems []OrderItem `json:"order_items"`
}

// 重写匿名方法
func (e Order) GetId() (id uint64) {
	return e.ID
}

type OrderItem struct {
	Entity
	ID        uint64  `json:"id"`
	OrderId   uint64  `json:"order_id"`
	Product   Product `json:"product"`
	Quantity  uint16  `json:"quantity"`
	UnitPrice float32 `json:"unit_price"`
}

// GetId 重写匿名方法
func (e OrderItem) GetId() (id uint64) {
	return e.ID
}
