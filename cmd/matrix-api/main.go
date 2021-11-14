package main

import (
	"flag"
	"github.com/toolkits/pkg/file"
	"go.uber.org/zap"
	"log"
	"matrix/internal/common/logger"
	"matrix/internal/matrix-api/conf"
	"matrix/internal/matrix-api/dao"
	"os"
	"os/signal"
	"syscall"

	matrixhttp "matrix/internal/matrix-api/http"
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

func main() {
	log.Println("start app!!!")

	app := matrixhttp.New()
	appConfig := conf.NewAppConfig(*config)
	if appConfig.Log == nil {
		zap.S().Fatal("log config error")
	}
	logger.Init(appConfig.Log)
	dao.Init(appConfig)

	go app.Run()
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
