package controllers

import (
	"Gozone/library/authorization"
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/models"
)

type UserController struct {
	BaseHandler
}

func (this *UserController) Register() {

	user := new(models.User)
	err := controller.ParseRequestStruct(this.Controller, &user)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	err, _ = new(dao.RegisterService).CheckRegister(user)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	err = new(dao.RegisterService).Register(user)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")

}

func (this *UserController) Login() {

	user := new(models.User)
	err := controller.ParseRequestStruct(this.Controller, &user)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	if user.UserName == "" && user.Email == "" {
		this.Response(1,"请填写用户名或邮箱")
	}
	if user.PassWord == "" {
		this.Response(1,"请填写登录密码")
	}

	UserDB, err := new(dao.LoginService).Login(user)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}

	token, err := new(dao.LoginService).CreateToken(UserDB)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	_ = authorization.AddUserToken(token, UserDB.Id)

	this.Response(enum.DefaultSuccess, "")
}

func (this *UserController) Logout() {
	this.DeleteCookie("admin-cookie")
	this.Redirect("/login", 302)
}
