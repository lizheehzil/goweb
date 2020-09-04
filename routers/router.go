package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{},"get:Get")
	//beego.AutoRouter(&controllers.UserController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.CMSController{})
}
