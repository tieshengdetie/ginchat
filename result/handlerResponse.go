package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WrappedHandlerFunc func(*gin.Context) interface{}

// WrapData 包装响应结果
func WrapData(handler WrappedHandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		resp := handler(c)
		// 响应 json
		c.JSON(http.StatusOK, resp)
	}
}
