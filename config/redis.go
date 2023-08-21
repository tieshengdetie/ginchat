package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var client *redis.Client

func InitRedisClient() {
	var ctx = context.Background()
	addr := viper.GetString("redis.addr")
	fmt.Println(addr)
	client = redis.NewClient(&redis.Options{
		Addr:        viper.GetString("redis.addr"),     // 连接地址
		Password:    viper.GetString("redis.password"), // 密码
		DB:          viper.GetInt("redis.db"),          // 数据库编号
		DialTimeout: 5 * time.Second,                   // 链接超时
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func GetRedisClient() *redis.Client {
	return client
}
