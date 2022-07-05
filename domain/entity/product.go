package entity

type Product struct {
	Entity
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	UnitPrice float32 `json:"unit_price"`
}

// GetId 重写匿名方法
func (e Product) GetId() (id uint64) {
	return e.ID
}
