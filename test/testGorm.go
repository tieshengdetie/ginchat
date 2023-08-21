package main

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(192.168.0.201:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	//db.AutoMigrate(&models.Relation{})
	//db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.UserBasic{})
	//user := &models.Message{}
	//db.AutoMigrate(user)
	//user.Name = "zhaojinsheng"
	//db.Create(user)
	//data := make([]*models.UserBasic, 10)
	//result := db.Find(&data)
	//user.Password = "nihao"
	//db.Save(user)
	//fmt.Println(result)
}
