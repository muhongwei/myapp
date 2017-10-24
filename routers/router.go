package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
}
