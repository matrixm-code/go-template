package http

import (
	"github.com/gin-gonic/gin"
	"gtemplate/internal/template-api/conf"
	"gtemplate/internal/template-api/http/controller"
)

type HttpServer struct {
	c  *conf.AppConfig
	v  *Validator
	cl *controller.Controller
	s  *gin.Engine
}

func NewHttpServer(c *conf.AppConfig, v *Validator, cl *controller.Controller) *HttpServer {
	return &HttpServer{c: c, v: v, cl: cl, s: gin.New()}
}

func (h HttpServer) Run() {
	h.s.Use(gin.Logger())
	h.s.Use(gin.Recovery())
	h.SetRouter(h.s)
	h.v.SetValidate()
	h.s.Run()
}
