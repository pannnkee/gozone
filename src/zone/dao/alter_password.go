package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"errors"
)

type AlterPasswordService struct {}

func (this *AlterPasswordService) Do(email, password, newPassword, repeatPassword string) (err error) {

	// 校验当前密码是否正确
	login := models.UserInstance.Login(email, password)
	if !login {
		return errors.New("密码错误,请重新输入")
	}

	// 判断两次输入密码是否一致
	if newPassword != repeatPassword {
		return errors.New("两次密码输入不一致,请检查")
	}

	// 更新密码
	user := new(models.User)
	user.Email = email
	user.PassWord = str.Md5(newPassword)
	_ = user.Updates()
	return nil
}
