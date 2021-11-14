package logic

import (
	"matrix/internal/matrix-api/dao"
	"matrix/internal/models"
)

func GetSampeList() []models.Sample {
	return dao.GetDao().GetSampleList()
}