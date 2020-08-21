package dao

import (
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"fmt"
)

type LoginService struct{}

func (this *LoginService) Login(userName, eMail, password string) error {

	user, err := new(models.User).GetPasswordByUserName(userName)
	if err != nil {
		return err
	}

	if user.PassWord  !=  str.Md5(password) {
		return fmt.Errorf("密码错误")
	}
	return nil
}
