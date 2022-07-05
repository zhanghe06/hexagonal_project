package entity

type Customer struct {
	Entity
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	BaseEntity
}

// GetId 重写匿名方法
func (e Customer) GetId() (id uint64) {
	return e.ID
}
