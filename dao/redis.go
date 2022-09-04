package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"rank/config"
)

var RedisClient *redis.Client

const redisDriver = "redis"

//NewRedisClient 初始化redis连接
func NewRedisClient() (*redis.Client, error) {
	// 获取配置
	redisConfig := config.Conf.Storages[redisDriver]
	// 初始化
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Open Redis Error:%v", err)
		return nil, err
	}
	return redisClient, nil
}

//GetRedisData 读取redis数据
func GetRedisData(key string) (interface{}, error) {
	//读取
	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		log.Errorf("read Redis Error:%v", err)
		return 0, err
	}
	return val, nil
}

// UpdateRedis 更新redisRating
func UpdateRedis(key string, val int) {

}
