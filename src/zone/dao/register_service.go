package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"errors"
	"time"
)

type RegisterService struct {}

func (this *RegisterService) Register(user *models.User) (err error) {
	user.CreatedTime = time.Now().Unix()
	user.PassWord = str.Md5(user.PassWord)
	err = user.Register()
	return
}


func (this *RegisterService) CheckRegister(user *models.User) (error, bool) {
	UsernameExist := new(models.User).UserNameExist(user.UserName)
	if UsernameExist {
		return errors.New("userName is exist"), false
	}

	eMailExist := new(models.User).UserNameExist(user.Email)
	if eMailExist {
		return errors.New("eMail is exist"), false
	}

	if user.PassWord != user.PassWord {
		return errors.New("password is nor equal"), false
	}

	return nil, true
}
