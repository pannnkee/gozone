package controllers

import (
	"Gozone/library/controller"
	"Gozone/library/enum"
	"Gozone/src/zone/dao"
	"Gozone/src/zone/model_view"
	"Gozone/src/zone/models"
	"fmt"
)

type UserController struct {
	BaseHandler
}

func (this *UserController) Register() {

	fmt.Println("收到请求")
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

	user := new(model_view.User)
	err := controller.ParseRequestStruct(this.Controller, &user)
	if err != nil {
		this.Response(enum.DefaultError,err.Error())
		return
	}

	if user.Email == "" {
		this.Response(1,"请填写登录邮箱")
	}
	if user.PassWord == "" {
		this.Response(1,"请填写登录密码")
	}

	err = new(dao.LoginService).Login(user.Email, user.PassWord)
	if err != nil {
		this.Response(enum.DefaultError, err.Error())
		return
	}
	this.Response(enum.DefaultSuccess, "")
}

func (this *UserController) Logout() {
	this.DeleteCookie("admin-cookie")
	this.Redirect("/login", 302)
}
