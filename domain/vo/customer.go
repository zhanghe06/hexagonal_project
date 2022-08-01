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
	Name              string                     `json:"name" binding:"required"` // 名称
	CustomerAddresses []CustomerAddressCreateReq `json:"addresses" binding:"omitempty,unique=Contact,dive"`
}

// CustomerUpdateReq .
type CustomerUpdateReq struct {
	Name string `json:"name,omitempty" binding:"omitempty"` // 名称
}

// CustomerGetListReq .
// form tag: ShouldBindQuery for api request
// json tag: json.Marshal/json.Unmarshal for map filter
type CustomerGetListReq struct {
	Limit     *int     `form:"limit,omitempty" json:"limit,omitempty" binding:"omitempty"`
	Offset    *int     `form:"offset,omitempty" json:"offset,omitempty" binding:"omitempty"`
	CreatedAt []string `form:"created_at[],omitempty" json:"created_at[],omitempty" binding:"omitempty"` // 创建时间
	Sorter    string   `form:"sorter,omitempty" json:"sorter,omitempty" binding:"omitempty"`             // 排序字段
}

type CustomerAddressCreateReq struct {
	Address   string `json:"address" binding:"required"`
	Contact   string `json:"contact" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"omitempty,email"`
	DefaultSt *uint8 `json:"default_st" binding:"required,eq=0|eq=1"`
}
