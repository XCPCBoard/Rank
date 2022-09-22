package router

import (
	"github.com/gin-gonic/gin"
	"rank/server"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", server.Ping)

	return r
}
