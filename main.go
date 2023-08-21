package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
)

func main() {

	r := router.Router()
	r.Run(":8081")
}
func init() {
	// 初始化数据库
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedisClient()
	//初始化验证器

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
}
