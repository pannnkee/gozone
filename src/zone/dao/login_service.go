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

func (this *LoginService) Login(eMail, password string) (cookie []byte, err error) {

	userInfo, err := new(models.User).GetUserInfo(eMail)
	// 登陆失败
	if err != nil || userInfo.Id < 1 {
		return nil, errors.New("账号或者密码错误")
	}
	if userInfo.Status == 1 {
		return nil, errors.New("该账号被禁止登陆")
	}

	if userInfo.PassWord != str.Md5(password) {
		return nil, errors.New("账号或者密码错误")
	}

	// 登陆成功
	userInfo.LoginTimes = userInfo.LoginTimes + 1
	userInfo.UpdateTime = time.Now().Unix()
	_ = userInfo.Updates()

	// 生成token
	token, err := new(LoginService).CreateToken(&userInfo)
	if err != nil {
		return nil, err
	}
	_ = authorization.AddUserToken(token, userInfo.Id)

	// 生成cookie
	m, err := json.Marshal(&userInfo)
	if err != nil {
		return nil, errors.New("生成cookie失败")
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
