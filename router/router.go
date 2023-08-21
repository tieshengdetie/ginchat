package router

import (
	"ginchat/controller"
	"ginchat/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	//swagger url
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//静态资源
	r.Static("/asset", "./asset")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	r.LoadHTMLGlob("views/**/*")
	//首页相关
	r.GET("/", controller.GetIndex)
	r.GET("/index", controller.GetIndex)
	r.GET("/toRegister", controller.ToRegister)
	//用户相关
	r.GET("/user/getUserList", controller.GetUserList)
	r.POST("/user/createUser", controller.CreateUser)
	r.POST("/user/updateUser", controller.UpdateUser)
	r.POST("/user/login", controller.Login)
	r.POST("/searchFriends", controller.SearchFriends)
	r.POST("/contact/addfriend", controller.AddFriends)
	//聊天
	r.GET("/toChat", controller.ToChat)
	//文件上传
	r.POST("/attach/upload", controller.Upload)
	r.GET("/ws", controller.SendMsg)
	return r
}
