package dao

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"matrix/internal/matrix-api/conf"
)


var dao *Dao

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{db: db}
}

func Init(c *conf.AppConfig) {
	db, err := gorm.Open(mysql.Open(c.Db.Addr), &gorm.Config{})
	if err != nil {
		zap.S().Fatal("get db error")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(c.Db.Idle)
	sqlDb.SetMaxOpenConns(c.Db.Max)

	dao = NewDao(db)
}

func GetDao() *Dao {
	return dao
}
