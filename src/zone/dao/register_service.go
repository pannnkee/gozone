package dao

import (
	"Gozone/library/util/str"
	"Gozone/library/verifycode"
	"Gozone/src/zone/models"
	"errors"
	"time"
)

type RegisterService struct{}


var (
	ErrUserNameExist      = errors.New("用户名已被注册")
	ErrEmailExist         = errors.New("邮件已被注册")
	ErrPasswordNotEqual   = errors.New("两次输入的密码不一致")
	ErrVerifyCodeNotRight = errors.New("输入的验证码错误, 请检查验证码")
	ErrVerifyCodeIsNil    = errors.New("验证码已失效或者无效，请重新发起验证请求")
)

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

	codeRedis := verifycode.Get(eMail)
	if codeRedis == "" {
		return ErrVerifyCodeIsNil, false
	}

	if code != codeRedis {
		return ErrVerifyCodeNotRight, false
	}

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
	verifycode.Del(eMail)
	return nil, true
}

