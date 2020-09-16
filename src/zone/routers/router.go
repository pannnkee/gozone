package routers

import (
	"Gozone/src/zone/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ZoneController{}, "*:Home")
	beego.Router("/login", &controllers.ZoneController{}, "*:Login")
	beego.Router("/register", &controllers.ZoneController{}, "*:Register")
	beego.Router("/profile", &controllers.ZoneController{}, "*:Profile")

	v1 := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/register", &controllers.UserController{}, "post,options:Register"),
			beego.NSRouter("/login", &controllers.UserController{}, "post,options:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "*:Logout"),
			beego.NSRouter("/alterPassword", &controllers.UserController{}, "*:AlterPassword"),
			beego.NSRouter("/alterAvatar", &controllers.UserController{},"*:AlterAvatar"),
		),
		beego.NSNamespace("/article",
			beego.NSRouter("/page/?type:int", &controllers.ArticleController{}, "*:PageList"),
			beego.NSRouter(":id:int", &controllers.ArticleController{}, "get:Get"),
		),
	)
	beego.AddNamespace(v1)
}
