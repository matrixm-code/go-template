package controller

import (
	"github.com/gin-gonic/gin"
	"matrix/internal/template-api/logic"
	"net/http"
)

type SampleController struct {
	logic *logic.SampleLogic
}

func NewSampleController(logic *logic.SampleLogic) *SampleController {
	return &SampleController{logic: logic}
}

func (s SampleController) GetAllSampleList(c *gin.Context) {
	result := s.logic.GetSampeList()
	c.JSON(http.StatusOK, gin.H{"err":"", "result": result})
}