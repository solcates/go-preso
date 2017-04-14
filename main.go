package main

import (
	_ "github.com/solcates/go-preso/routers"
	"github.com/astaxie/beego"
)


func init() {
	orm.RegisterDataBase("default", "mongo", "", 30)

}

func main() {
	beego.Run()
}

