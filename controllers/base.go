package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare(){
	uname := c.GetSession("uname")
	if uname == nil {
		log.Info("User not authenticated")
		c.Ctx.Output.SetStatus(401)
		c.ServeJSON()
	} else {
		log.Info("User is authenticated")
	}
}
