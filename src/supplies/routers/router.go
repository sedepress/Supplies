package routers

import (
	"supplies/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{}, "get:Index")
    beego.AutoRouter(&controllers.AdminController{})
}
