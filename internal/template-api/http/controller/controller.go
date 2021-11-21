package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	SampleController *SampleController
}

func NewController(
	sampleController *SampleController,

) *Controller {
	return &Controller{
		SampleController: sampleController,
	}
}

func (reciver Controller) Ping(c *gin.Context) {
	c.String(http.StatusOK,"pong")
}