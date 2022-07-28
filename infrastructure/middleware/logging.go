package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		skipPaths := []string{"/health/ready", "/health/alive"}
		for _, v := range skipPaths {
			if v == c.Request.URL.Path {
				return
			}
		}

		serverName := "hexagonal_project"

		// 开始时间
		timeStart := time.Now()

		// 若存在，则传递，便于链路追踪
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Writer.Header().Set("X-Request-Id", requestId)

		// 获取请求
		requestBodyCopy := new(bytes.Buffer)
		// Read the whole body
		_, err := io.Copy(requestBodyCopy, c.Request.Body)
		if err != nil {
			return // 退出中间件
		}
		requestBodyBytes := requestBodyCopy.Bytes()
		// Replace the body with a reader that reads from the buffer
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBodyBytes))

		var requestBody string
		requestBody = string(requestBodyBytes) // 默认原始

		requestBodyDecoder := json.NewDecoder(requestBodyCopy)
		var paramsMap map[string]interface{}
		if err = requestBodyDecoder.Decode(&paramsMap); err == nil {
			// 数据脱敏
			// delete(paramsMap, "password")
			if _, ok := paramsMap["password"]; ok {
				paramsMap["password"] = "******"
			}

			paramsByte, e := json.Marshal(paramsMap)
			if e == nil {
				requestBody = string(paramsByte) // 反向解析
			}
		}

		//writer := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		writer := &CustomResponseWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
		c.Writer = writer

		// Context 赋值
		c.Set("request_id", requestId) // 链路核心字段
		c.Set("server_name", serverName)

		c.Next()

		// 结束时间
		timeEnd := time.Now()
		// 执行时间
		latency := timeEnd.Sub(timeStart)

		logInfo := make(map[string]interface{})
		logInfo["req_id"] = requestId
		logInfo["service_name"] = serverName
		logInfo["time_start"] = timeStart
		logInfo["time_end"] = timeEnd
		logInfo["latency"] = latency
		logInfo["client_ip"] = c.ClientIP()
		logInfo["req_host"] = c.Request.Host
		logInfo["req_method"] = c.Request.Method
		logInfo["req_path"] = c.Request.URL.Path
		logInfo["req_query"] = c.Request.URL.RawQuery
		logInfo["req_header"] = ""
		logInfo["req_body"] = ""
		logInfo["res_status_code"] = c.Writer.Status()
		logInfo["res_error_msg"] = c.Errors.String()
		logInfo["res_body"] = ""

		if c.Writer.Status()/100 > 2 {
			logInfo["req_header"] = c.Request.Header
			logInfo["req_body"] = requestBody
			logInfo["res_body"] = writer.body.String()
		}

		logInfoByte, _ := json.Marshal(logInfo)
		fmt.Println(string(logInfoByte))
	}
}
