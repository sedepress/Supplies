package main

import (
	"supplies/controllers"
	_ "supplies/routers"
	"github.com/astaxie/beego"
	"supplies/utils"
)

func main() {
	utils.InitSql()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

