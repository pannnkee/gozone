package controllers


type LoginController struct {
	BaseHandler
}

func (this *LoginController) Login() {
	name := this.GetString("user_name")
	password := this.GetString("password")
	email := this.GetString("email")

	if name == "" || email == "" {
		this.Response(1,"请填写用户名或邮箱")
	}
	if password == "" {
		this.Response(1,"请填写登录密码")
	}


}

func (this *LoginController) Logout() {
	this.DeleteCookie("admin-cookie")
	this.Redirect("/login", 302)
}
