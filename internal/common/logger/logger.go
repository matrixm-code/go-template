package logger

import (
	"go.uber.org/zap"
	"log"
	"matrix/internal/models"
)

/*
日志配置格式:
level: debug
encode: json
outputPaths:
  - stdout
encoderConfig:
  - stderr
*/

var (
	logger *zap.Logger
	Suger  *zap.SugaredLogger
)

func init() {
	logger, _ = zap.NewProduction()
	Suger = logger.Sugar()

}

func Init(c *models.LogConfig) {
	lvl := zap.NewAtomicLevel()
	lvl.UnmarshalText([]byte(c.Level))
	config := zap.Config{
		Level:       lvl,
		Development: c.Development,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		DisableCaller:     false,
		DisableStacktrace: false,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       c.OutputPaths,
		ErrorOutputPaths:  c.ErrorOutputPaths,
	}
	lg, err := config.Build()
	if err != nil {
		log.Println(err)
		log.Fatalln("create logger failed")
	}
	zap.ReplaceGlobals(lg)
}
