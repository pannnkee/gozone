package models

import (
	"Gozone/library/conn"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id             int64    `gorm:"column:id" json:"id"`
	UserName       string `gorm:"column:user_name" json:"user_name"`
	PassWord       string `gorm:"column:password" json:"password"`
	RepeatPassword string `gorm:"-" json:"repeat_password"`
	Email          string `gorm:"column:email" json:"email"`
	Status         int    `gorm:"column:status" json:"status"`
	Mobile         string `gorm:"column:mobile" json:"mobile"`
	CreatedTime    int64  `gorm:"column:created_time" json:"created_time"`
	UpdateTime     int64  `gorm:"column:update_time" json:"update_time"`
}

func (this *User) TableName() string {
	return "user"
}

func (this *User) Register() (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Create(&this).Error
	return
}

func (this *User) GetPasswordByUserName(userName string) (password string, err error) {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("user_name=?", userName).First(&user).Error
	password = user.PassWord
	return
}

func (this *User) GetPasswordByEmail(eMail string) (password string, err error) {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("email=?", eMail).First(&user).Error
	password = user.PassWord
	return
}

func (this *User) UserNameExist(userName string) bool {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("user_name=?", userName).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (this *User) EmailExist(email string) bool {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("email=?", email).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (this *User) GetUserByUserId(id int64) (*User, error) {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("id=?", id).First(&user).Error
	return &user, err
}

func (this *User) GetUserByUserName(userName string) (*User, error) {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("user_name=?", userName).First(&user).Error
	return &user, err
}

func (this *User) GetUserByEmail(eMail string) (*User, error) {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("email=?", eMail).First(&user).Error
	return &user, err
}
