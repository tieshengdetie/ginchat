package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db          *gorm.DB
	RedisClient *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
func InitMysql() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢sql阈值
		LogLevel:      logger.Info, //级别
		Colorful:      true,
	})
	dns := viper.GetString("mysql.dns")
	Db, _ = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
}

func InitRedisClient() {
	var ctx = context.Background()
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        addr,            // 连接地址
		Password:    password,        // 密码
		DB:          db,              // 数据库编号
		DialTimeout: 5 * time.Second, // 链接超时
	})
	pong, err := RedisClient.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
