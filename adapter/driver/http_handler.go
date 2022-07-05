package driver

import (
	"github.com/gin-gonic/gin"
)

// HttpHandlerInterface .
type HttpHandlerInterface interface {
	// RegisterAPI 注册API
	RegisterAPI(engine *gin.Engine)
}
