package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	netHttp "net/http"
	"strings"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 1.8以前的版本不支持泛型，需要自行实现
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func formatHeader(header netHttp.Header) string {
	keyMaxLong := 0
	for k, _ := range header {
		keyMaxLong = maxInt(keyMaxLong, len(k))
	}
	var headerList []string
	for name, headers := range header {
		name = strings.ToLower(name)
		for _, h := range headers {
			format := fmt.Sprintf("%%%ds: %%s", keyMaxLong)
			headerList = append(headerList, fmt.Sprintf(format, name, h))
		}
	}
	return strings.Join(headerList, "\n")
}

func RequestContentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBodyBytes []byte
		if c.Request.Body != nil {
			reqBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// 仅仅记录失败请求的内容
		if c.Writer.Status()/100 == 2 {
			return
		}

		// Content-Type: application/json 重新序列化
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			var reqBodyMap map[string]interface{}
			if err := json.Unmarshal(reqBodyBytes, &reqBodyMap); err == nil {
				// 数据脱敏 TODO
				if _, ok := reqBodyMap["password"]; ok {
					reqBodyMap["password"] = "******"
				}
				reqBodyBytes, _ = json.MarshalIndent(reqBodyMap, "", "  ")
			}
		}

		var resBodyBytes []byte
		resBodyBytes = blw.body.Bytes()
		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			var resBodyMap map[string]interface{}
			if err := json.Unmarshal(resBodyBytes, &resBodyMap); err == nil {
				// 数据脱敏 TODO
				if _, ok := resBodyMap["password"]; ok {
					resBodyMap["password"] = "******"
				}
				resBodyBytes, _ = json.MarshalIndent(resBodyMap, "", "  ")
			}
		}

		c.Set("req_body", string(reqBodyBytes))
		c.Set("res_body", string(resBodyBytes))
		c.Set("req_header", formatHeader(c.Request.Header))
		c.Set("res_header", formatHeader(c.Writer.Header()))
	}
}
