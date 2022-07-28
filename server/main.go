package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexagonal_project/adapter/driver"
	"hexagonal_project/adapter/driver/http"
	"hexagonal_project/infrastructure/config"
	"hexagonal_project/infrastructure/middleware"
	"log"
	"os"
	"strings"
	"time"
)

type server struct {
	// Rest API
	healthRestHandler   driver.HttpHandlerInterface
	customerRestHandler driver.HttpHandlerInterface
	orderRestHandler    driver.HttpHandlerInterface
}


func (s *server) Start() {
	conf := config.NewConfig()
	go func() {
		engine := gin.New()
		//engine.Use(gin.Recovery())

		formatter := func(p gin.LogFormatterParams) string {

			errMsg := strings.TrimSpace(p.ErrorMessage)
			logInfo := fmt.Sprintf("[logger] %s %s %s %s %s %s %d %s \"%s\" \"%s\"\n",
				p.TimeStamp.Format(time.RFC3339),
				p.Keys["request_id"],
				p.ClientIP,
				p.Method,
				p.Path,
				p.Request.Proto,
				p.StatusCode,
				p.Latency,
				p.Request.UserAgent(),
				errMsg,
			)
			if errMsg != "" {
				logInfo += fmt.Sprintf("[req_header]\n%s\n", p.Keys["req_header"])
				logInfo += fmt.Sprintf("[req_body]\n%s\n", p.Keys["req_body"])
				logInfo += fmt.Sprintf("[res_header]\n%s\n", p.Keys["res_header"])
				logInfo += fmt.Sprintf("[res_body]\n%s\n", p.Keys["res_body"])
			}
			return logInfo
		}
		logConf := gin.LoggerConfig{
			SkipPaths: []string{"/health/ready", "/health/alive"},
			Output:    os.Stderr,
			Formatter: formatter,
		}
		engine.Use(
			//gin.Logger(),
			gin.LoggerWithConfig(logConf),
			middleware.RequestIdMiddleware(),
			middleware.RequestContentMiddleware(),
			middleware.RecoveryMiddleware(),
			middleware.TokenAuthMiddleware(),
		)
		engine.UseRawPath = true

		// 注册API
		s.healthRestHandler.RegisterAPI(engine)
		s.customerRestHandler.RegisterAPI(engine)
		s.orderRestHandler.RegisterAPI(engine)

		url := fmt.Sprintf(
			"%s:%d",
			conf.Server.Host,
			conf.Server.Port,
		)
		if err := engine.Run(url); err != nil {
			log.Fatal(err)
		}
	}()
}

func main() {
	s := &server{
		healthRestHandler:   http.NewHealthHttpHandler(),
		customerRestHandler: http.NewCustomerHttpHandler(),
		orderRestHandler:    http.NewOrderHttpHandler(),
	}
	s.Start()

	select {}
}
