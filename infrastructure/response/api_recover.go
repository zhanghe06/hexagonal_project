package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

func ApiRecover(c *gin.Context) {
	if rec := recover(); rec != nil {

		// 获取一段堆栈
		//stackBuf := make([]byte, 2048)
		//stackBufLen := runtime.Stack(stackBuf[:], false)
		//errorStack := string(stackBuf[:stackBufLen])

		// 获取全部堆栈
		c.Set("level", "panic")
		c.Set("error_stack", string(Stack()))

		err := fmt.Errorf("%v", rec)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}
