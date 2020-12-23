package model_view

// 用户
type User struct {
	Id             int64  `json:"id"`				// 用户id
	UserName       string `json:"username"`			// 用户姓名
	PassWord       string `json:"password"`			// 密码
	RepeatPassword string `json:"repeat_password"`	// 重复密码
	Email          string `json:"email"`			// 邮箱
	NewPassword    string `json:"new_password"`		// 新密码
	VerifyCode     string `json:"verify_code"`		// 验证码
}
