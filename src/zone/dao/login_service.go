package dao

import (
	"Gozone/library/authorization"
	"Gozone/library/jwt"
	"Gozone/library/model"
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"errors"
)

type LoginService struct{}

func (this *LoginService) Login(eMail, password string) (err error) {

	userInfo, err := new(models.User).GetUserInfo(eMail)
	if err != nil {
		return err
	}

	if userInfo.PassWord != str.Md5(password) {
		return errors.New("账号或者密码错误")
	}

	token, err := new(LoginService).CreateToken(&userInfo)
	if err != nil {
		return err
	}
	_ = authorization.AddUserToken(token, userInfo.Id)
	return
}

func (this *LoginService) CreateToken(user *models.User) (string, error) {

	userToken := new(model.UserToken)
	if user.Id > 0 {
		userToken.Id = user.Id
	}
	if user.UserName != "" {
		userToken.UserName = user.UserName
	}
	if user.PassWord != "" {
		userToken.Password = user.PassWord
	}
	token, err := new(jwt.ZoneJsonWebTokenHelper).CreateToken(userToken)
	if err != nil {
		return "", errors.New("create token failed:" + err.Error())
	}
	return token, nil
}
