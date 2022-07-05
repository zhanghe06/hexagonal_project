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

		//formatter := func(p gin.LogFormatterParams) string {
		//	return fmt.Sprintf("[logger] %s %s %s %d %s\n",
		//		p.TimeStamp.Format("2006-01-02_15:04:05"),
		//		p.Path,
		//		p.Method,
		//		p.StatusCode,
		//		p.ClientIP,
		//	)
		//}
		logConf := gin.LoggerConfig{
			SkipPaths: []string{"/health/ready", "/health/alive"},
			Output:    os.Stderr,
			//Formatter: formatter,
		}
		engine.Use(
			//gin.Logger(),
			gin.LoggerWithConfig(logConf),
			middleware.RecoveryMiddleware(),
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
