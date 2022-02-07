package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:dbctime" json:"dbctime"`
	UpdatedAt time.Time      `gorm:"column:dbutime" json:"dbutime"`
	DeletedAt gorm.DeletedAt `gorm:"column:dbdtime" json:"dbdtime" gorm:"index"`
}
