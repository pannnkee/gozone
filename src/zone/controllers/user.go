package controllers

import (
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/src/zone/auth"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/model_view"
	"Gozone/src/zone/models"
	"time"
)

type UserController struct {
	BaseHandler
}

func (this *UserController) Register() {

	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	err, _ = new(dao.RegisterService).Do(modelUser.UserName, modelUser.Email, modelUser.PassWord, modelUser.RepeatPassword)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, this.GetString("next"))

}

func (this *UserController) Login() {
	//TODO 检查账号密码合法性
	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	if modelUser.Email == "" {
		this.Response(1,"请填写登录邮箱")
	}
	if modelUser.PassWord == "" {
		this.Response(1,"请填写登录密码")
	}

	m, err := new(dao.LoginService).Do(modelUser.Email, modelUser.PassWord)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	now := time.Now().Unix()
	log := models.Log{
		Ip:           this.GetIP(),
		UserID:       modelUser.Id,
		UserName:     modelUser.UserName,
		LoginTime:    now,
		LoginTimeStr: time.Unix(now, 0).Format("2006-01-02 15:04:05"),
	}
	_ = log.AddLoginLog()

	userInfo, _ := models.UserInstance.UserInfo(modelUser.Email)
	if userInfo.Avatar == "" {
		userInfo.Avatar = "/static/img/default_avatar.png"
	}
	this.SetCK(auth.ZoneToken, string(m), 168)
	this.SetSession(SESSION_USER_KEY, userInfo)
	this.Response(0, this.GetString("next"))
}

func (this *UserController) Logout() {

	this.MustLogin()
	this.DeleteCookie(auth.ZoneToken)
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}

func (this *UserController) AlterPassword() {
	var modelUser model_view.User
	err := controller.ParseRequestStruct(this.Controller, &modelUser)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}
	err = new(dao.AlterPasswordService).Do(this.User.Email, modelUser.PassWord, modelUser.NewPassword, modelUser.RepeatPassword)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")
	return
}

func (this *UserController) AlterData() {
	this.Response(enum.DefaultError,"暂时未开放此功能")
}
