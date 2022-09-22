package server

import (
	"github.com/gin-gonic/gin"
	"rank/sort"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func UpdateRating(c *gin.Context) {
	sort.Flush()
	users := AllUserRating()
	c.JSON(200, gin.H{
		"user": users,
	})
}
