package service

import (
	"gozone/library/util"
	"gozone/src/zone/dao"
	"gozone/src/zone/models"
	"time"
)

type RegisterService struct{}

// 注册账号
// @param userName 用户名
// @param eMail 邮件
// @param password 密码
// @param repeatPassword 重复密码
// @param code 验证码
// @return 错误信息 是否注册成功
func (this *RegisterService) Do(userName, eMail, password, repeatPassword, code string) (error, bool) {

	UsernameExist := dao.UserInstance.UserNameExist(userName)
	if UsernameExist {
		return ErrUserNameExist, false
	}

	eMailExist := dao.UserInstance.EmailExist(eMail)
	if eMailExist {
		return ErrEmailExist, false
	}

	if password != repeatPassword {
		return ErrPasswordNotEqual, false
	}

	//codeRedis := verifycode.Get(eMail)
	//if codeRedis == "" {
	//	return ErrVerifyCodeIsNil, false
	//}
	//
	//if code != codeRedis {
	//	return ErrVerifyCodeNotRight, false
	//}

	user := models.User{
		UserName:    userName,
		Email:       eMail,
		PassWord:    util.Md5(password),
		CreatedTime: time.Now().Unix(),
	}
	err := dao.UserInstance.Create(&user)
	if err != nil {
		return err, false
	}

	// 注册成功删除验证码
	//verifycode.Del(eMail)
	return nil, true
}

