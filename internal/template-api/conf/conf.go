package conf

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/toolkits/pkg/file"
	"go.uber.org/zap"
	"log"
	"gtemplate/internal/models"
)

type AppConfig struct {
	Db  *models.DbConfig  `yaml:"db"`
	Log *models.LogConfig `yaml:"log"`
}


func NewAppConfig(cPath string) *AppConfig {
	bs, err := file.ReadBytes(cPath)
	if err != nil {
		log.Fatalf("cannot read yml[%s]: %v\n", cPath, err)
		return nil
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalf("cannot read yaml [%s]: %v", cPath, err)
	}

	viper.SetDefault("log", map[string]interface{}{
		"level": "debug",
		"deployment": false,
		"encoding": "json",
		"outputPaths": []string{"stdout"},
		"errorOutputPaths": []string{"stderr"},
	})

	viper.SetDefault("db", map[string]interface{}{
		"max": 2000,
		"idle": 500,
		"debug": false,
	})

	var config = new(AppConfig)
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
		zap.S().Error(err)
	}

	return config
}


