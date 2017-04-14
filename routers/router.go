package routers

import (
	"github.com/solcates/go-preso/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/presos", &controllers.MainController{})
}
