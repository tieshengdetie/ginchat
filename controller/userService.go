package controller

import (
	"ginchat/models"
	"ginchat/result"
	"ginchat/server/ws"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList
// @Description 首页
// @Tags 用户相关接口
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {

	p := utils.NewPagination(c)
	data := models.GetUserList(p)
	c.JSON(http.StatusOK, result.OK.WithData(gin.H{
		"page":  p.Page,
		"limit": p.Size,
		"total": p.Total,
		"user":  data,
	}))
}

// CreateUser
// @Description
// @Tags 用户相关接口
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	//接收参数
	user := models.UserBasic{}
	if err := c.ShouldBind(&user); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	rePassword := c.PostForm("repassword")
	if user.Name == "" || password == "" {
		c.JSON(http.StatusOK, result.ErrUserNameOrPwd)
		return
	}
	//校验用户名
	isName := models.FindUserByName(user.Name)

	if isName.Name != "" {
		c.JSON(http.StatusOK, result.ErrUserExsit)
		return
	}
	//校验密码
	if password != rePassword {
		c.JSON(http.StatusOK, result.ErrUserRePwd)
		return
	}
	//生成随机salt
	salt := utils.RandStr(20)
	user.Salt = salt
	user.Password = utils.MakePassword(password, salt)
	models.CreateUser(user)

	c.JSON(http.StatusOK, result.OK)
}

// UpdateUser
// @Description
// @Tags 用户相关接口
// @param id formData string false "id"
// @param name formData string false "用户名"
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	//接收参数
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	avatar := c.PostForm("icon")
	user.Name = c.PostForm("name")
	user.Avatar = avatar
	models.UpdateUser(user)

	c.JSON(http.StatusOK, result.OK)
}

// Login
// @Description
// @Tags 用户相关接口
// @param password formData string false "密码"
// @param name formData string false "用户名"
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/login [post]

func Login(c *gin.Context) {

	//定义字段验证结构体
	type paramStruct struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	paramUser := paramStruct{}
	if err := c.ShouldBind(&paramUser); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, result.ErrUserName)
		return
	}
	//校验密码
	isPwd := utils.ValidPassword(password, user.Salt, user.Password)
	if !isPwd {
		c.JSON(http.StatusOK, result.ErrUserPwd)
		return
	}
	//生成token
	tokenClaims := utils.Users{ID: user.ID, Username: user.Name}
	token, _ := utils.GenToken(tokenClaims)
	c.JSON(http.StatusOK, result.OK.WithData(gin.H{"ID": user.ID, "token": token}))
}

func SearchFriends(c *gin.Context) {
	//接收参数
	p := utils.NewPagination(c)
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	userData := models.SearcFriends(uint(userId), p)
	c.JSON(http.StatusOK, result.OK.WithData(gin.H{
		"page":  p.Page,
		"limit": p.Size,
		"total": p.Total,
		"user":  userData,
	}))
}

func AddFriends(c *gin.Context) {
	//接收参数
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	targetName := c.PostForm("targetName")
	//验证targetId是否存在
	targetUser := models.FindUserByName(targetName)
	if targetUser.ID == 0 {
		c.JSON(http.StatusOK, result.Err.WithMsg("此用户不存在"))
		return
	}
	if uint(userId) == targetUser.ID {
		c.JSON(http.StatusOK, result.Err.WithMsg("自己不能添加自己哦!"))
		return
	}
	//验证是否已经添加过此好友
	isFriend := models.FindRelationByUserIdAndTargetId(uint(userId), targetUser.ID)
	if isFriend.ID != 0 {
		c.JSON(http.StatusOK, result.Err.WithMsg("您已经添加过此好友"))
		return
	}
	isOk := models.AddFriends(uint(userId), targetUser.ID)
	if !isOk {
		c.JSON(http.StatusOK, result.Err.WithMsg("添加失败"))
		return
	}
	c.JSON(http.StatusOK, result.OK)
}

//防止跨域站点的伪造请求

func SendMsg(c *gin.Context) {

	ws.Chat(c.Writer, c.Request)
	//var upgrader = websocket.Upgrader{
	//	ReadBufferSize:  1024,
	//	WriteBufferSize: 1024,
	//	//设置允许跨域
	//	CheckOrigin: func(r *http.Request) bool {
	//		return true
	//	},
	//	//设置请求协议
	//	Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	//}
	//ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// 处理WebSocket消息
	//for {
	//	messageType, p, err := ws.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//
	//	fmt.Println("messageType:", messageType)
	//	fmt.Println("p:", string(p))
	//
	//	// 输出WebSocket消息内容
	//	write, err := c.Writer.Write(p)
	//	if err != nil {
	//		fmt.Println(write, err)
	//	}
	//}
	//
	//// 关闭WebSocket连接
	//defer func(ws *websocket.Conn) {
	//	err := ws.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}(ws)
}
