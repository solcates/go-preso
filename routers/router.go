package routers

import (
	"github.com/solcates/go-preso/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/logout", &controllers.LoginController{}, "post:Logout;get:Logout")
	beego.Router("/api/presos", &controllers.PresoController{})
	beego.Router("/api/me", &controllers.LoginController{}, "get:Me;post:Me")

}
