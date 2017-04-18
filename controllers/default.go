package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "presos.thalesignite.com"
	c.Data["Email"] = "scates@vormetric.com"
	c.TplName = "index.html"
	c.Render()
}
