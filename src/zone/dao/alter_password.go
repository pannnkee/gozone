package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
)

type AlterPasswordService struct{}

func (this *AlterPasswordService) Do(email, password, newPassword, repeatPassword string) (err error) {

	// 校验当前密码是否正确
	login := models.UserInstance.Login(email, password)
	if !login {
		return ErrAccountOrPassword
	}

	// 判断两次输入密码是否一致
	if newPassword != repeatPassword {
		return ErrPasswordNotEqual
	}

	// 更新密码
	user := new(models.User)
	user.Email = email
	user.PassWord = str.Md5(newPassword)

	exmap := map[string]interface{}{
		"email":    email,
		"password": str.Md5(newPassword),
	}
	_ = user.Updates(email, exmap)
	return nil
}
