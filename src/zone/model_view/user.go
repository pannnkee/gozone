package model_view

type User struct {
	Id             int64  `json:"id"`
	UserName       string `json:"username"`
	PassWord       string `json:"password"`
	RepeatPassword string `json:"repeatpassword"`
	Email          string `json:"email"`
	NewPassword    string `json:"newpassword"`
	VerifyCode     string `json:"verifycode"`
}
