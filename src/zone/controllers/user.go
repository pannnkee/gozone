package controllers

import (
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/library/util"
	"Gozone/src/zone/auth"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/model_view"
	"Gozone/src/zone/models"
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

	err, _ = new(dao.RegisterService).Register(modelUser.UserName, modelUser.Email, modelUser.PassWord, modelUser.RepeatPassword)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")

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

	m, err := new(dao.LoginService).Login(modelUser.Email, modelUser.PassWord)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	userInfo, _ := models.UserInstance.UserInfo(modelUser.Email)
	this.SetCK(auth.ZoneToken, string(m), 168)
	this.SetSession(SESSION_USER_KEY, userInfo)

	userInfoMap, _ := util.Struct2JsonMap(userInfo)
	this.Data["User"] = userInfoMap
	this.Response(0, "登录成功")
}

func (this *UserController) Logout() {

	this.MustLogin()
	this.DeleteCookie(auth.ZoneToken)
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}
