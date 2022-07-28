package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hexagonal_project/infrastructure/response"
	"net/http"
	"strings"
)

type UserInfo struct {
	ID   uint64 `json:"user_id"`
	Name string `json:"user_name"`
}

// In [1]: import json
// In [2]: import base64
// In [3]: base64.b64encode(json.dumps({"user_id": 1, "user_name": "admin"}).encode("utf-8"))
// Out[3]: b'eyJ1c2VyX2lkIjogMSwgInVzZXJfbmFtZSI6ICJhZG1pbiJ9'

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			err := response.NewApiError(
				"token required",
				response.Unauthorized,
			)
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		//if token != os.Getenv("API_TOKEN") {
		//	err := response.NewApiError(
		//		"token invalid",
		//		response.Unauthorized,
		//	)
		//	_ = c.AbortWithError(http.StatusUnauthorized, err)
		//	return
		//}

		tokenObject, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			err = response.NewApiError(
				err.Error(),
				response.Unauthorized,
			)
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		var userInfo *UserInfo
		if err = json.Unmarshal(tokenObject, &userInfo); err != nil {
			err = response.NewApiError(
				err.Error(),
				response.Unauthorized,
			)
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user_id", userInfo.ID)
		c.Set("user_name", userInfo.Name)

		c.Next()
	}
}
