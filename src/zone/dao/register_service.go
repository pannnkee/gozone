package dao

import "Gozone/src/zone/models"

type RegisterService struct {}

func (this *RegisterService) Register(Username, password, repeatPassword, eMail string) (err error) {

	UsernameExist := new(models.User).UserNameExist(Username)
	if UsernameExist {

	}

	eMailExist := new(models.User).UserNameExist(eMail)
	if eMailExist {

	}

	if password != repeatPassword {

	}
	return nil
}
