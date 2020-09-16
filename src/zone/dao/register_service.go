package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"errors"
	"time"
)

type RegisterService struct {}

func (this *RegisterService) Do(userName, eMail, password, repeatPassword string) (error, bool) {

	UsernameExist := models.UserInstance.UserNameExist(userName)
	if UsernameExist {
		return errors.New("userName is exist"), false
	}

	eMailExist := models.UserInstance.EmailExist(eMail)
	if eMailExist {
		return errors.New("eMail is exist"), false
	}

	if password != repeatPassword {
		return errors.New("password is not equal"), false
	}

	user := models.User{
		UserName: userName,
		Email: eMail,
		PassWord: str.Md5(password),
		CreatedTime: time.Now().Unix(),
	}
	err := user.Register()
	if err != nil {
		return err, false
	}
	return nil, true
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

	if user.PassWord != user.RepeatPassword {
		return errors.New("password is nor equal"), false
	}

	return nil, true
}
