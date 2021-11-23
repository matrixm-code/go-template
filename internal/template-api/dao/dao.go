package dao

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gtemplate/internal/template-api/conf"
)


var dao *Dao

type Dao struct {
	conf *conf.AppConfig
	db *gorm.DB
}

func NewDao(c *conf.AppConfig) *Dao {
	db, err := gorm.Open(mysql.Open(c.Db.Addr), &gorm.Config{})
	if err != nil {
		zap.S().Fatal("get db error")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(c.Db.Idle)
	sqlDb.SetMaxOpenConns(c.Db.Max)
	return &Dao{conf: c, db: db}
}
