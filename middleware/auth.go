package middleware

import (
	"ginchat/result"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, result.ErrUnauthorized.WithMsg("没有携带token"))
			c.Abort() //结束后续操作
			return
		}
		log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, result.ErrUnauthorized.WithMsg("token格式有误"))
			c.Abort()
			return
		}

		//解析token包含的信息
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, result.ErrUnauthorized.WithMsg("无效的token"))
			c.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		c.Set("userinfo", claims)
		c.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}
