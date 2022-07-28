package vo

type UriCustomerIdReq struct {
	CustomerId uint64 `uri:"customer_id" binding:"required"`
}

type UriCustomerAddressIdReq struct {
	CustomerId uint64 `uri:"customer_id" binding:"required"`
	AddressId  uint64 `uri:"address_id" binding:"required"`
}

// CustomerCreateReq .
type CustomerCreateReq struct {
	Name              string                     `binding:"required" form:"name" json:"name"` // 名称
	CustomerAddresses []CustomerAddressCreateReq `binding:"omitempty,dive" form:"addresses" json:"addresses"`
}

// CustomerUpdateReq .
type CustomerUpdateReq struct {
	Name string `binding:"omitempty" json:"name,omitempty"` // 名称
}

// CustomerGetListReq .
type CustomerGetListReq struct {
	Limit     *int     `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset    *int     `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	CreatedAt []string `binding:"omitempty" form:"created_at[],omitempty" json:"created_at[],omitempty"` // 创建时间
	Sorter    string   `binding:"omitempty" form:"sorter,omitempty" json:"sorter,omitempty"`             // 排序字段
}

type CustomerAddressCreateReq struct {
	Address   string `binding:"required" form:"address" json:"address"`
	Contact   string `binding:"required" form:"contact" json:"contact"`
	Phone     string `binding:"required" form:"phone" json:"phone"`
	Email     string `binding:"omitempty,email" form:"email" json:"email"`
	DefaultSt *uint8 `binding:"required" form:"default_st" json:"default_st"`
}
