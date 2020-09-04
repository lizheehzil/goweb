package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["goweb/controllers:CMSController"] = append(beego.GlobalControllerRouter["goweb/controllers:CMSController"],
		beego.ControllerComments{
			Method:           "AllBlock",
			Router:           "/all",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["goweb/controllers:CMSController"] = append(beego.GlobalControllerRouter["goweb/controllers:CMSController"],
		beego.ControllerComments{
			Method:           "StaticBlock",
			Router:           "/staticblock/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["goweb/controllers:UserController"] = append(beego.GlobalControllerRouter["goweb/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Testautorouter",
			Router:           "/testautorouter",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
