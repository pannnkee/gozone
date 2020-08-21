package models

import (
	"Gozone/library/conn"
	"github.com/jinzhu/gorm"
)


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

func (this *User) GetPasswordByUserName(userName string) (user *User, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("username=?", userName).Find(&user).Error
	return
}

func (this *User) UserNameExist(userName string) bool {
	var user User
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("username=?", userName).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (this *User) EmailExist(email string) bool {
	var user User
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("email=?", email).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
