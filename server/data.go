package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"rank/dao"
	"rank/model"
)

func AllUserRating() []model.KV {
	user, err := dao.RedisClient.SMembers(context.Background(), "rating").Result()
	if err != nil {
		panic(err)
	}
	gxuRating := []model.KV{}
	for _, u := range user {
		val, er := dao.RedisClient.ZScore(context.Background(), "rating", u).Result()
		if er != nil {
			val = 0
			log.Errorf("get %s rating error:%v", u, er)
		}
		gxuRating = append(gxuRating, model.KV{u, val})
	}
	return gxuRating
}

