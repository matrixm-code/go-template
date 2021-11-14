package http

import (
	"github.com/gin-gonic/gin"
	"matrix/internal/matrix-api/http/controller"
)


func SetRouter(r *gin.Engine) {

	noAuth := r.Group("/")
	{
		noAuth.GET("/ping", controller.HealthCheck)
		noAuth.GET("/sample", controller.SampleList)

	}

}
