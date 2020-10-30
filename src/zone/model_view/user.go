package model_view

// 用户
type User struct {
	Id             int64  `json:"id"`
	UserName       string `json:"username"`
	PassWord       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Email          string `json:"email"`
	NewPassword    string `json:"new_password"`
	VerifyCode     string `json:"verify_code"`
}
