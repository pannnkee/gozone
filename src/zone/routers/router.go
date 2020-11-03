package routers

import (
	"gozone/src/zone/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ZoneController{}, "*:Home")
	beego.Router("/login", &controllers.ZoneController{}, "*:Login")
	beego.Router("/register", &controllers.ZoneController{}, "*:Register")
	beego.Router("/profile", &controllers.ZoneController{}, "*:Profile")
	beego.Router("/alterPassword", &controllers.ZoneController{}, "*:AlterPassword")
	beego.Router("/alterData", &controllers.ZoneController{}, "*:AlterData")
	beego.Router("/timeline", &controllers.ZoneController{}, "*:TimeLine")
	beego.Router("/about", &controllers.ZoneController{}, "*:About")
	beego.Router("/archive", &controllers.ZoneController{}, "*:Archive")

	v1 := beego.NewNamespace("/v1/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/register/*", &controllers.UserController{}, "post,options:Register"),
			beego.NSRouter("/login/*", &controllers.UserController{}, "post,options:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "*:Logout"),
			beego.NSRouter("/alterPassword", &controllers.UserController{}, "post,options:AlterPassword"),
			beego.NSRouter("/alterData", &controllers.UserController{}, "*:AlterData"),
			beego.NSRouter("/verifyCode", &controllers.UserController{},"post,options:VerifyCode"),
		),
		beego.NSNamespace("/article",
			beego.NSRouter("/page/?type:int", &controllers.ArticleController{}, "*:PageList"),
			beego.NSRouter(":id:int", &controllers.ArticleController{}, "get:Get"),
			beego.NSRouter(":id:int", &controllers.ArticleController{}, "post,options:Comment"),
		),

		beego.NSNamespace("/static",
			beego.NSRouter("/search", &controllers.StaticController{}, "*:Search"),
			),
	)
	beego.AddNamespace(v1)
}
