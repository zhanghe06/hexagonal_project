package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"hexagonal_project/adapter/driver"
	"hexagonal_project/domain/entity"
	"hexagonal_project/domain/service"
	"hexagonal_project/domain/vo"
	"hexagonal_project/infrastructure/response"
	"hexagonal_project/port/service_port"
	"net/http"
	"strings"
	"sync"
)

type customerHttpHandler struct {
	customerService service_port.CustomerServicePort
	customerAddressService service_port.CustomerAddressServicePort
}

var (
	customerHttpOnce sync.Once
	customerHttpHand driver.HttpHandlerInterface
)

func NewCustomerHttpHandler() driver.HttpHandlerInterface {
	customerHttpOnce.Do(func() {
		customerHttpHand = &customerHttpHandler{
			customerService: service.NewCustomerService(),
			customerAddressService: service.NewCustomerAddressService(),
		}
	})
	return customerHttpHand
}

func (h *customerHttpHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/customer/:customer_id", h.getCustomerHandler)
	engine.POST("/customer", h.createCustomerHandler)
	engine.PUT("/customer/:customer_id", h.updateCustomerHandler)
	engine.GET("/customer", h.getCustomerListHandler)
	engine.DELETE("/customer/:customer_id", h.deleteCustomerHandler)
	engine.GET("/customer/:customer_id/address/:address_id", h.getCustomerAddressHandler)
	//engine.POST("/customer/:customer_id/address", h.createCustomerAddressHandler)
	//engine.PUT("/customer/:customer_id/address/:address_id", h.updateCustomerAddressHandler)
	//engine.GET("/customer/:customer_id/address", h.getCustomerAddressListHandler)
	//engine.DELETE("/customer/:customer_id/address/:address_id", h.deleteCustomerAddressHandler)
}

func (h *customerHttpHandler) getCustomerHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriCustomerIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	res, err := h.customerService.GetInfo(c, uriIdReq.CustomerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 原始错误
			//_ = c.AbortWithError(http.StatusNotFound, err)

			// 封装错误
			apiErr := response.NewApiError(
				err.Error(),
				response.CustomerNotFound,
			)
			_ = c.AbortWithError(http.StatusNotFound, apiErr)
			return
		}
		// 异常错误
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, res)
}

func (h *customerHttpHandler) createCustomerHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var customerCreateReq vo.CustomerCreateReq
	if err := c.ShouldBindJSON(&customerCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var customerEntity entity.Customer
	customerEntity.Name = customerCreateReq.Name

	customerEntity.CustomerAddresses = make([]entity.CustomerAddress, 0)

	for _, v := range customerCreateReq.CustomerAddresses {
		customerAddress := entity.CustomerAddress{
			Address: v.Address,
			Contact: v.Contact,
			Phone: v.Phone,
			Email: v.Email,
			DefaultSt: *v.DefaultSt,
		}
		customerEntity.CustomerAddresses = append(customerEntity.CustomerAddresses, customerAddress)
	}

	// 逻辑处理
	res, err := h.customerService.Create(c, customerEntity)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Header("Location", c.FullPath()+fmt.Sprintf("/%d", res.ID))
	c.JSON(http.StatusCreated, gin.H{
		"id": res.GetId(),
	})
}

func (h *customerHttpHandler) updateCustomerHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriCustomerIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var customerUpdateReq vo.CustomerUpdateReq
	if err := c.ShouldBindJSON(&customerUpdateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var data map[string]interface{}
	reqBytes, err := json.Marshal(customerUpdateReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(reqBytes, &data)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err = h.customerService.Update(c, uriIdReq.CustomerId, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 原始错误
			//_ = c.AbortWithError(http.StatusNotFound, err)

			// 封装错误
			apiErr := response.NewApiError(
				err.Error(),
				response.CustomerNotFound,
			)
			_ = c.AbortWithError(http.StatusNotFound, apiErr)
			return
		}
		// 异常错误
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *customerHttpHandler) getCustomerListHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var customerGetListReq vo.CustomerGetListReq
	if err := c.ShouldBindQuery(&customerGetListReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	customerGetListReqBytes, err := json.Marshal(customerGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(customerGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 扩展参数（范围查询）
	var filterQuery []string
	filterArgs := make([]interface{}, 0)
	if v, ok := filter["created_at[]"]; ok {
		filterQuery = append(filterQuery, "created_at BETWEEN ? AND ?")
		filterArgs = append(filterArgs, v.([]interface{})...)
		delete(filter, "created_at[]")
	}
	filterQueries := strings.Join(filterQuery, " AND ")
	args := make([]interface{}, 0)
	if filterQueries != "" {
		args = append(args, filterQueries)
		args = append(args, filterArgs...)
	}

	// 逻辑处理
	total, data, err := h.customerService.GetList(c, filter, args...)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, gin.H{
		"count": total,
		"data":  data,
	})
}

func (h *customerHttpHandler) deleteCustomerHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriCustomerIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err := h.customerService.Delete(c, uriIdReq.CustomerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 原始错误
			//_ = c.AbortWithError(http.StatusNotFound, err)

			// 封装错误
			apiErr := response.NewApiError(
				err.Error(),
				response.CustomerNotFound,
			)
			_ = c.AbortWithError(http.StatusNotFound, apiErr)
			return
		}
		// 异常错误
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}

func (h *customerHttpHandler) getCustomerAddressHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriCustomerAddressIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	res, err := h.customerAddressService.GetInfo(c, uriIdReq.CustomerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 原始错误
			//_ = c.AbortWithError(http.StatusNotFound, err)

			// 封装错误
			apiErr := response.NewApiError(
				err.Error(),
				response.CustomerNotFound,
			)
			_ = c.AbortWithError(http.StatusNotFound, apiErr)
			return
		}
		// 异常错误
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, res)
}
