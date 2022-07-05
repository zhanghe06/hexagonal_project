package errors

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func (e *ApiError) Error() string {
	es, _ := json.Marshal(e)
	return string(es)
}

var _ error = &ApiError{}

var (
	//i18n         = make(map[int]map[string]string)
	c2m = make(map[int]string)
)

// NewApiError 新建一个ApiError
func NewApiError(cause string, code int) *ApiError {
	return &ApiError{
		Cause:   cause,
		Code:    code,
		Message: c2m[code],
	}
}

// --------
// 错误码 分组自定义错误码，每组使用const分别定义
// --------

// 通用异常
const (
	BadRequest          = 400000000
	Unauthorized        = 401000000
	Forbidden           = 403000000
	NotFound            = 404000000
	Conflict            = 409000000
	InternalServerError = 500000000
)

const (
	// Customer 客户模块
	Customer           = 300000
	CustomerBadRequest = BadRequest + Customer + iota
	CustomerRepeated
	CustomerDisabled
	CustomerNotFound = NotFound + Customer + iota
)

func init() {
	// 通用异常
	c2m[BadRequest] = http.StatusText(http.StatusBadRequest)
	c2m[Unauthorized] = http.StatusText(http.StatusUnauthorized)
	c2m[Forbidden] = http.StatusText(http.StatusForbidden)
	c2m[NotFound] = http.StatusText(http.StatusNotFound)
	c2m[Conflict] = http.StatusText(http.StatusConflict)
	c2m[InternalServerError] = http.StatusText(http.StatusInternalServerError)

	// Customer 客户模块
	c2m[CustomerBadRequest] = "Customer " + http.StatusText(http.StatusBadRequest)
	c2m[CustomerRepeated] = "Customer Repeated"
	c2m[CustomerDisabled] = "Customer Disabled"
	c2m[CustomerNotFound] = "Customer " + http.StatusText(http.StatusNotFound)

	//fmt.Println(c2m)
}
