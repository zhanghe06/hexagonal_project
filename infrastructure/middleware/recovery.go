package middleware

import (
	"github.com/gin-gonic/gin"
	"hexagonal_project/infrastructure/errors"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 错误处理
		defer func() {
			for _, err := range c.Errors {
				statusCode := c.Writer.Status()

				switch interface{}(err.Err).(type) {
				// 接口错误
				case *errors.ApiError:
					c.AbortWithStatusJSON(statusCode, err.Err)
				// 组件错误 todo
				// 系统异常
				default:
					if statusCode >= http.StatusInternalServerError {
						// 服务器的内部错误
						c.AbortWithStatusJSON(statusCode, gin.H{
							"code":    errors.InternalServerError,
							"message": http.StatusText(http.StatusInternalServerError),
							"cause":   err.Error(),
						})
						return
					} else {
						// 未封装的请求错误
						c.AbortWithStatusJSON(statusCode, gin.H{
							"code":    statusCode * 1000000,
							"message": http.StatusText(statusCode),
							"cause":   err.Error(),
						})
						return
					}
				}
				return
			}
		}()

		//contentType := c.ContentType()
		//if contentType != "" {
		//	c.Writer.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", c.ContentType()))
		//}
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
