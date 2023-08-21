package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

var clusterClient *redis.ClusterClient

func init() {
	// 连接redis集群
	ctx := context.Background()
	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{ // 填写master主机
			"192.168.0.200:6379",
			"192.168.0.201:6379",
			"192.168.0.202:6379",
		},
		Password:     "123",                 // 设置密码
		DialTimeout:  50 * time.Microsecond, // 设置连接超时
		ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	// 发送一个ping命令,测试是否通
	s := clusterClient.Do(ctx, "ping").String()
	fmt.Println(s)
}
func main() {
	ctx := context.Background()
	err := clusterClient.Set(ctx, "k1", "tiesheng", 0).Err()
	if err != nil {

		panic(err)
	}
	getValue := clusterClient.Get(ctx, "k1")
	fmt.Println(getValue.Val(), getValue.Err())
}
