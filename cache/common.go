package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
)

func ConnRedis() {
	rd := redis.NewClient(&redis.Options{
		Addr: "", // url
		Password: "",
		DB:1,   // 1号数据库
	})
	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Println("ping err :",err)
		return
	}
	fmt.Println(result)
}
