package result

// 错误码规则:
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//              ----------------------------------------------------------
//                  第1位               2、3位                  4、5位
//              ----------------------------------------------------------
//                服务级错误码          模块级错误码	         具体错误码
//              ----------------------------------------------------------

var (
	// OK
	OK  = response(200, "ok") // 通用成功
	Err = response(500, "")   // 通用错误

	// ErrParam 服务级错误码
	ErrParam     = response(10001, "参数错误")
	ErrSignParam = response(10002, "签名参数错误")

	// ErrUserService 模块级错误码 - 用户模块
	ErrUserService   = response(20100, "用户服务异常")
	ErrUserPhone     = response(20101, "用户手机号不合法")
	ErrUserCaptcha   = response(20102, "用户验证码有误")
	ErrUserName      = response(20103, "用户名不存在")
	ErrUserPwd       = response(20104, "登录密码不正确")
	ErrUserNameOrPwd = response(20105, "用户名或密码不能为空")
	ErrUserExsit     = response(20106, "用户名已经注册")
	ErrUserRePwd     = response(20107, "两次密码不一致")
	// ErrOrderService 库存模块
	ErrOrderService = response(20200, "订单服务异常")
	ErrOrderOutTime = response(20201, "订单超时")

	//文件上传
	ErrFileUpload     = response(50001, "获取上传文件失败")
	ErrFileUploadFail = response(50002, "文件上传失败")
	// ErrLoginUser  登录模块错误
	ErrLoginUser    = response(40001, "token 异常")
	ErrUnauthorized = response(40001, "token 认证失败")
)
