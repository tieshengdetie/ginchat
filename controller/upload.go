package controller

import (
	"ginchat/result"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	request := c.Request
	upFile, head, err := request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, result.ErrFileUpload)
	}

	oriName := head.Filename
	tem := strings.Split(oriName, ".")
	subfix := ".png"

	if len(tem) > 1 {
		subfix = "." + tem[len(tem)-1]
	}
	filename := utils.Md5Encode(utils.RandStr(6)+string(time.Now().Unix())) + subfix
	//filename := fmt.Sprintf("%d%04v%d", time.Now().Unix(), rand.Int(), subfix)
	dirFile := "./asset/upload/" + filename
	dstFile, err := os.Create(dirFile)
	_, err = io.Copy(dstFile, upFile)

	if err != nil {
		c.JSON(http.StatusOK, result.ErrFileUploadFail)
	}
	c.JSON(http.StatusOK, result.OK.WithData(gin.H{"url": dirFile}))
}
