package routers

import (
	"Gozone/src/zone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

    beego.Router("/register", )
    beego.Router("/login", &controllers.LoginController{}, "*:Login")
}
