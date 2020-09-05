package controllers

type ZoneController struct {
	BaseHandler
}

func (this *ZoneController) Home() {
	this.TplName = "base.html"
}

func (this *ZoneController) Login() {
	this.TplName = "login.html"
}

func (this *ZoneController) Register() {
	this.TplName = "register.html"
}

//
//func (this *ZoneController) Article() {
//
//}