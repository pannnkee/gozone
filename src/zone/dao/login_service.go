package dao

import (
	"Gozone/library/authorization"
	"Gozone/library/jwt"
	"Gozone/library/model"
	"Gozone/library/util/str"
	"Gozone/src/zone/models"
	"encoding/json"
	"errors"
	"time"
)

type LoginService struct{}


func (this *LoginService) Do(eMail, password string) (cookie []byte, err error) {

	userInfo, err := new(models.User).UserInfo(eMail)
	// 登陆失败
	if err != nil || userInfo.Id < 1 {
		return nil, ErrAccountOrPassword
	}
	if userInfo.Status == 1 {
		return nil, ErrAccountNotAllowed
	}

	if userInfo.PassWord != str.Md5(password) {
		return nil, ErrAccountOrPassword
	}

	// 登陆成功
	exmap := map[string]interface{} {
		"login_time" : userInfo.LoginTimes + 1,
		"update_time" : time.Now().Unix(),
	}
	_ = userInfo.Updates(userInfo.Email, exmap)

	// 生成token
	token, err := new(LoginService).CreateToken(&userInfo)
	if err != nil {
		return nil, err
	}
	_ = authorization.AddUserToken(token, userInfo.Id)

	// 生成cookie
	m, err := json.Marshal(&userInfo)
	if err != nil {
		return nil, ErrServerInternal
	}
	return m, nil
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
