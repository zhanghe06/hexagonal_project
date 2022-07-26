package entity

type Customer struct {
	Entity            `json:"-"`
	ID                uint64            `json:"id"`
	Name              string            `json:"name"`
	CustomerAddresses []CustomerAddress `json:"addresses,omitempty"` // 详情显示，列表隐藏
	BaseEntity
}

// GetId 重写匿名方法
func (e Customer) GetId() (id uint64) {
	return e.ID
}

type CustomerAddress struct {
	Entity    `json:"-"`
	ID        uint64 `json:"id"`
	Address   string `json:"address"`
	Contact   string `json:"contact"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	DefaultSt uint8  `json:"default_st"`
	BaseEntity
}

// GetId 重写匿名方法
func (e CustomerAddress) GetId() (id uint64) {
	return e.ID
}
