package dao

import "errors"

var (
	ErrAccountOrPassword = errors.New("账号或者密码错误")
	ErrAccountNotAllowed = errors.New("该账号被禁止登陆")
	ErrServerInternal = errors.New("服务器内部错误")

	ErrUserNameExist      = errors.New("用户名已被注册")
	ErrEmailExist         = errors.New("邮件已被注册")
	ErrPasswordNotEqual   = errors.New("两次输入的密码不一致")
	ErrVerifyCodeNotRight = errors.New("输入的验证码错误, 请检查验证码")
	ErrVerifyCodeIsNil    = errors.New("验证码已失效或者无效，请重新发起验证请求")
)