package main

import (
	"flag"
	"github.com/toolkits/pkg/file"
	"go.uber.org/zap"
	"gtemplate/internal/common/logger"
	"gtemplate/internal/template-api/conf"
	"gtemplate/internal/template-api/dao"
	gtemplatehttp "gtemplate/internal/template-api/http"
	"gtemplate/internal/template-api/http/controller"
	"gtemplate/internal/template-api/logic"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	config *string
)

func init() {
	config = flag.String("c", "configs/conf.yaml", "config")
	flag.Parse()
	if *config == "" || !file.IsExist(*config) {
		zap.S().Info("config file not exist!!")
		os.Exit(0)
	}
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   lizx
// @contact.email  lizx@yuanfudao.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	log.Println("start app!!!")

	appConfig := conf.NewAppConfig(*config)
	if appConfig.Log == nil {
		zap.S().Fatal("log config error")
	}
	logger.Init(appConfig.Log)

	// 初始化dao
	serviceDao := dao.NewDao(appConfig)

	// 初始化validate
	validator := gtemplatehttp.NewValidator(appConfig)

	// 初始化proxy (默认不添加)

	// 初始化logic
	sampleLogic := logic.NewSampleLogic(serviceDao)

	// 初始化controller
	sampleController := controller.NewSampleController(sampleLogic)
	controller := controller.NewController(
		sampleController,
	)

	s := gtemplatehttp.NewHttpServer(appConfig, validator, controller)

	go s.Run()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sig
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//todo
			zap.S().Info("shutdown")
			return
		case syscall.SIGHUP:
		default:
			return

		}
	}
}
