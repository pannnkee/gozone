package models

import (
	"Gozone/library/conn"
	"Gozone/library/util/str"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id             int64  `gorm:"column:id" json:"id"`
	UserName       string `gorm:"column:user_name" json:"user_name"`
	Email          string `gorm:"column:email" json:"email"`
	Mobile         string `gorm:"column:mobile" json:"mobile"`
	PassWord       string `gorm:"column:password" json:"password"`
	RepeatPassword string `gorm:"-" json:"repeat_password"`
	Status         int64  `gorm:"column:status" json:"status"`
	LoginTimes     int64  `gorm:"column:login_time" json:"login_time"`
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

func (this *User) UserNameExist(userName string) bool {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("user_name=?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (this *User) EmailExist(eMail string) bool {
	user := User{}
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err := db.Where("email=?", eMail).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (this *User) Login(eMail, password string) (login bool) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	db = db.Where("email=?", eMail).Where("password=?", str.Md5(password))

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return false
	}
	if count != 1 {
		return false
	}
	return true
}

func (this *User) UserInfo(eMail string) (user User, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(this)
	err = db.Where("email=?", eMail).First(&user).Error
	return
}

func (this *User) Updates() error {
	db := conn.GetORMByName("zone")
	return db.Save(&this).Error
}
