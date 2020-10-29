package dao

import (
	"Gozone/library/conn"
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"github.com/jinzhu/gorm"
)

type UserDao struct {}

// 新建一个用户信息
// @return err 错误信息
func (this *UserDao) Create(user *models.User) (err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err = db.Create(&this).Error
	return
}

// 用户名是否存在
// @param userName 用户名
// @return 是否存在
func (this *UserDao) UserNameExist(userName string) bool {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err := db.Where("user_name=?", userName).First(models.UserInstance).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// eMail是否已经注册
// @param eMail 邮件
// @return 是否已经注册
func (this *UserDao) EmailExist(eMail string) bool {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err := db.Where("email=?", eMail).First(models.UserInstance).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// 登录是否成功
// @param eMail 邮件
// @param password 密码
// @return login 是否登录
func (this *UserDao) Login(eMail, password string) (login bool) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
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

// 用户详情
// @param eMail 用户邮件
// @return user 用户信息
// @return err 错误信息
func (this *UserDao) UserInfo(eMail string) (user models.User, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err = db.Where("email=?", eMail).First(&user).Error
	return
}

// 根据email更新用户信息
// @param email 邮件
// @param exMap 更新条目
// @return err 错误信息
func (this *UserDao) Updates(email string, exMap map[string]interface{}) error {
	db := conn.GetORMByName("zone")
	return db.Model(models.UserInstance).Where("email=?", email).Updates(exMap).Error
}

// 获取所有用户信息
// @return data 所有用户信息
// @return err 错误信息
func (this *UserDao) GetAll() (data []*models.User, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err = db.Order("id asc").Find(&data).Error
	return
}

// 根据UserId获取用户详情
// @param UserId 用户ID
// @param data 用户详情
// @param err 错误信息
func (this *UserDao) Get(UserId int64) (data models.User, err error) {
	db := conn.GetORMByName("zone")
	db = db.Model(models.UserInstance)
	err = db.Where("id=?", UserId).Take(&data).Error
	return
}
