package controllers

type HomeController struct {
	BaseHandler
}

func (this *HomeController) Home() {



	this.TplName = "base.html"
}

func (this *HomeController) Login() {
	this.TplName = "login.html"
}

func (this *HomeController) Register() {
	this.TplName = "register.html"
}


func (this *HomeController) Article() {
	this.TplName = "article.html"
}