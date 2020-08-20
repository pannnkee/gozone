package models

import "Gozone/library/conn"

type User struct {
	Id          int    `gorm:"column:id" json:"id"`
	UserName    string `gorm:"column:user_name" json:"user_name"`
	PassWord    string `gorm:"column:password" json:"password"`
	Email       string `gorm:"column:email" json:"email"`
	Status      int    `gorm:"column:status" json:"status"`
	Mobile      string `gorm:"column:mobile" json:"mobile"`
	CreatedTime int64  `gorm:"column:created_time" json:"created_time"`
	UpdateTime  int64  `gorm:"column:update_time" json:"update_time"`
}

func (this *User) TableName() string {
	return "user"
}

func (this *User) Login(username, password string) (err error) {
	db := conn.GetORMByName("zone")
	db.Model(this)
}
