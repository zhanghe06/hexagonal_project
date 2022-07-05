package http

import (
	"github.com/gin-gonic/gin"
	"hexagonal_project/adapter/driver"
	"net/http"
	"sync"
)

type healthHttpHandler struct {
}

var (
	healthHttpOnce sync.Once
	healthHttpHand driver.HttpHandlerInterface
)

func NewHealthHttpHandler() driver.HttpHandlerInterface {
	healthHttpOnce.Do(func() {
		healthHttpHand = &healthHttpHandler{}
	})
	return healthHttpHand
}

// RegisterAPI 注册API
func (h *healthHttpHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/health/ready", h.getHealth)
	engine.GET("/health/alive", h.getAlive)
}

func (h *healthHttpHandler) getHealth(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "ready")
}

func (h *healthHttpHandler) getAlive(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "alive")
}
