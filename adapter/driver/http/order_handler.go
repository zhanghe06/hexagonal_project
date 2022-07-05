package http

import (
	"github.com/gin-gonic/gin"
	"hexagonal_project/adapter/driver"
	"hexagonal_project/domain/service"
	"hexagonal_project/port/service_port"
	"net/http"
	"sync"
)

type orderHttpHandler struct {
	orderService service_port.OrderServicePort
}

var (
	orderHttpOnce sync.Once
	orderHttpHand driver.HttpHandlerInterface
)

func NewOrderHttpHandler() driver.HttpHandlerInterface {
	orderHttpOnce.Do(func() {
		orderHttpHand = &orderHttpHandler{
			orderService: service.NewOrderService(),
		}
	})
	return orderHttpHand
}

func (h *orderHttpHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/order/:id", h.getInfoHandler)
	engine.POST("/order", h.createHandler)
	engine.POST("/order/:id", h.updateHandler)
	engine.GET("/order", h.getListHandler)
	engine.DELETE("/order/:id", h.deleteHandler)
}

func (h *orderHttpHandler) getInfoHandler(c *gin.Context) {
	// 响应处理
	c.JSON(http.StatusOK, nil)
}

func (h *orderHttpHandler) createHandler(c *gin.Context) {
	// 响应处理
	c.JSON(http.StatusOK, nil)
}

func (h *orderHttpHandler) updateHandler(c *gin.Context) {
	// 响应处理
	c.JSON(http.StatusOK, nil)
}

func (h *orderHttpHandler) getListHandler(c *gin.Context) {
	// 响应处理
	c.JSON(http.StatusOK, nil)
}

func (h *orderHttpHandler) deleteHandler(c *gin.Context) {
	// 响应处理
	c.JSON(http.StatusOK, nil)
}
