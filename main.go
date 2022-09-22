package main

import (
	"github.com/gin-gonic/gin"
	"rank/dao"
	"rank/server"
)

// 主入口函数
func main() {
	r := gin.Default()
	r.POST("/Board", server.UpdateRating)
	r.Run()
}

func init() {
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	dbClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
	dao.DBClient = dbClient
}
