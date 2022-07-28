package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 若存在，则传递，便于链路追踪
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set("request_id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)

		c.Next()
	}
}
