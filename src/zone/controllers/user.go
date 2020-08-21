package controllers

import (
	"Gozone/library/enum"
	"Gozone/src/zone/dao"
)

type UserController struct {
	BaseHandler
}

func (this *UserController) Register() {
	Username := this.GetString("user_name")
	password := this.GetString("password")
	repeatPassword := this.GetString("repeat_password")
	eMail := this.GetString("email")

	new(dao.RegisterService).Register(Username, password, repeatPassword, eMail)

}

func (this *UserController) Login() {
	Username := this.GetString("user_name")
	password := this.GetString("password")
	eMail := this.GetString("email")

	if Username == "" && eMail == "" {
		this.Response(1,"请填写用户名或邮箱")
	}
	if password == "" {
		this.Response(1,"请填写登录密码")
	}

	err := new(dao.LoginService).Login(Username, eMail, password)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")
	return
}

func (this *UserController) Logout() {
	this.DeleteCookie("admin-cookie")
	this.Redirect("/login", 302)
}
