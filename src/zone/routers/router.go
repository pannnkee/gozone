package routers

import (
	"Gozone/src/zone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

    beego.Router("/register", &controllers.UserController{}, "*:Register")
    beego.Router("/login", &controllers.UserController{}, "*:Login")
    beego.Router("/logout", &controllers.UserController{}, "*:Logout")

}
