package router

import (
	"backend/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", health)
	r.POST("/find-pairs", handler.FindPairHandler)

	return r
}

func health(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "project go is up",
	})
}
