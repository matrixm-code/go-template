package controller

import (
	"github.com/gin-gonic/gin"
	"gtemplate/internal/template-api/logic"
	"net/http"
)

type SampleController struct {
	logic *logic.SampleLogic
}

func NewSampleController(logic *logic.SampleLogic) *SampleController {
	return &SampleController{logic: logic}
}

// GetAllSampleList godoc
// @Summary sample
// @Description 一个简单的GET 接口
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} models.Sample
// @Router /sample [get]
func (s SampleController) GetAllSampleList(c *gin.Context) {
	result := s.logic.GetSampeList()
	c.JSON(http.StatusOK, gin.H{"err": "", "result": result})
}
