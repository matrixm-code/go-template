package http

import (
	"github.com/gin-gonic/gin"
)


func New() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	//g.Use(middleware.ZapLogger())
	//g.Use(middleware.ZapRecovery(true))

	SetRouter(g)

	return g
}