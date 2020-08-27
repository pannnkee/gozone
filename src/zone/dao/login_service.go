package dao

import (
	"Gozone/library/jwt"
	"Gozone/library/model"
	"Gozone/src/zone/models"
	"errors"
)

type LoginService struct{}

func (this *LoginService) Login(userName, eMail, password string) () {
	new(models.User).
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
