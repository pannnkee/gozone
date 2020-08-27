package model_view

type User struct {
	Id             int64  `json:"id"`
	UserName       string `json:"user_name"`
	PassWord       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Email          string `json:"email"`
}
