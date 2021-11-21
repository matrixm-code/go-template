package http

import (
	"github.com/gin-gonic/gin"
)


func (h *HttpServer) SetRouter(r *gin.Engine) {

	noAuth := r.Group("/")
	{
		noAuth.GET("/sample", h.cl.SampleController.GetAllSampleList)
		noAuth.GET("/ping", h.cl.Ping)
	}

}