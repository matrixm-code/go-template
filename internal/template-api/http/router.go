package http

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gtemplate/docs"
)

func (h *HttpServer) SetRouter(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	noAuth := r.Group("/")
	{
		noAuth.GET("/sample", h.cl.SampleController.GetAllSampleList)
		noAuth.GET("/ping", h.cl.Ping)
	}

}