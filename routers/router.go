package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	//beego.Router("/echoService",&controllers.MainControllers{},"get:EchoService")
	beego.Router("/echoService",&controllers.MainController{},"get:EchoService")
}
