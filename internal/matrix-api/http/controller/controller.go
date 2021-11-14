package controller

import (
	"github.com/gin-gonic/gin"
	"matrix/internal/matrix-api/logic"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func SampleList(c *gin.Context) {
	result := logic.GetSampeList()
	c.JSON(http.StatusOK, gin.H{"result": result})
}