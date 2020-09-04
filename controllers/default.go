package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("This is get method!")
}

func (this *UserController) Post() {
	this.Ctx.WriteString("This is post method!")
}

//@router /testautorouter [get]
func (this *UserController) Testautorouter() {
	this.Ctx.WriteString("This is TestAutoRouter method!")
}

//@router /testhot [get]
func (this *UserController) TestHot() {
	this.Ctx.WriteString("This is TestHot method!")

}
