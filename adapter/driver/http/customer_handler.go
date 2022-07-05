package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
}

var (
	customerHttpOnce sync.Once
	customerHttpHand driver.HttpHandlerInterface
)

func NewCustomerHttpHandler() driver.HttpHandlerInterface {
	customerHttpOnce.Do(func() {
		customerHttpHand = &customerHttpHandler{
			customerService: service.NewCustomerService(),
		}
	})
	return customerHttpHand
}

func (h *customerHttpHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/customer/:id", h.getInfoHandler)
	engine.POST("/customer", h.createHandler)
	engine.POST("/customer/:id", h.updateHandler)
	engine.GET("/customer", h.getListHandler)
	engine.DELETE("/customer/:id", h.deleteHandler)
}

func (h *customerHttpHandler) getInfoHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	res, err := h.customerService.GetInfo(c, uriIdReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, res)
}

func (h *customerHttpHandler) createHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var customerCreateReq vo.CustomerCreateReq
	if err := c.ShouldBindJSON(&customerCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO 测试Context 后续放在认证中间件处理
	c.Set("userId", "admin")

	// 参数转换
	var customerEntity entity.Customer
	customerEntity.Name = customerCreateReq.Name

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

func (h *customerHttpHandler) updateHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriIdReq
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
	err = h.customerService.Update(c, uriIdReq.ID, data)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *customerHttpHandler) getListHandler(c *gin.Context) {
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

func (h *customerHttpHandler) deleteHandler(c *gin.Context) {
	// 异常捕获
	defer response.ApiRecover(c)

	// 请求处理
	var uriIdReq vo.UriIdReq
	if err := c.ShouldBindUri(&uriIdReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err := h.customerService.Delete(c, uriIdReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}
