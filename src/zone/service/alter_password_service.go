package service

import (
	"gozone/library/util"
	"gozone/src/zone/dao"
	"gozone/src/zone/models"
)

type AlterPasswordService struct{}

// 修改密码服务
// @param email 邮件
// @param password 密码
// @param newPassword 新密码
// @param repeatPassword 重复新密码
// @return err 错误信息
func (this *AlterPasswordService) Do(email, password, newPassword, repeatPassword string) (err error) {

	// 校验当前密码是否正确
	login := dao.UserInstance.Login(email, password)

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
	user.PassWord = util.Md5(newPassword)

	exmap := map[string]interface{}{
		"email":    email,
		"password": util.Md5(newPassword),
	}
	_ = dao.UserInstance.Updates(email, exmap)
	return nil
}
