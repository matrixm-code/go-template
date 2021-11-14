package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID      uint           `json:"id" gorm:"primarykey"`
	Dbctime time.Time      `json:"dbctime"`
	Dbutime time.Time      `json:"dbutime"`
	Dbdtime gorm.DeletedAt `json:"dbdtime" gorm:"index"`
}
