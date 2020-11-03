package model

type UserToken struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Status   int64  `json:"status"`
}
