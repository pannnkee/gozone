package model_view

type User struct {
	Id             int64  `json:"id"`
	UserName       string `json:"username"`
	PassWord       string `json:"password"`
	RepeatPassword string `json:"repeatpassword"`
	Email          string `json:"email"`
}
