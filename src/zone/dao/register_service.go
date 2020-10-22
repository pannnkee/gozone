package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"time"
)

type RegisterService struct{}

func (this *RegisterService) Do(userName, eMail, password, repeatPassword, code string) (error, bool) {

	UsernameExist := models.UserInstance.UserNameExist(userName)
	if UsernameExist {
		return ErrUserNameExist, false
	}

	eMailExist := models.UserInstance.EmailExist(eMail)
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
		PassWord:    str.Md5(password),
		CreatedTime: time.Now().Unix(),
	}
	err := user.Register()
	if err != nil {
		return err, false
	}

	// 注册成功删除验证码
	//verifycode.Del(eMail)
	return nil, true
}

