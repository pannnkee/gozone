package dao

import (
	"Gozone/library/jwt"
	"Gozone/library/model"
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"errors"
	"fmt"
)

type LoginService struct{}

func (this *LoginService) Login(user *models.User) (*models.User,error) {

	var password string
	var err error
	var UserDB *models.User

	if user.UserName != "" && user.Email == "" {
		password, err = new(models.User).GetPasswordByUserName(user.UserName)
		UserDB, _ = new(models.User).GetUserByUserName(user.UserName)
	}
	if user.Email != "" && user.UserName == ""{
		password, err = new(models.User).GetPasswordByEmail(user.Email)
		UserDB, _ = new(models.User).GetUserByEmail(user.Email)
	}
	if err != nil {
		return nil,err
	}

	if password != str.Md5(user.PassWord) {
		return nil,fmt.Errorf("密码错误")
	}
	return UserDB,nil
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
