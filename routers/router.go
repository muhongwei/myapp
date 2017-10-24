package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
}
