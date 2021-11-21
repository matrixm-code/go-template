package dao

import (
	"matrix/internal/models"
)

func (d *Dao)GetSampleList() (samples []models.Sample) {
	d.db.Find(&samples)
	return
}
