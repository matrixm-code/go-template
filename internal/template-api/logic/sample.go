package logic

import (
	"matrix/internal/models"
	"matrix/internal/template-api/dao"
)

type SampleLogic struct {
	dao *dao.Dao
}

func NewSampleLogic(dao *dao.Dao) *SampleLogic {
	return &SampleLogic{dao: dao}
}

func (s *SampleLogic) GetSampeList() []models.Sample {
	return s.dao.GetSampleList()
}
