package controller

import (
	"ginchat/models"
	"ginchat/result"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var user = models.User{}

	if err := c.ShouldBind(&user); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}
	//查询手机号有没有注册
	mobileExsit := models.FindUserByMobile(user.Mobile)
	if mobileExsit.ID > 0 {
		c.JSON(http.StatusOK, result.OK.WithMsg("该手机号已经注册"))
		return
	}
}
