package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/user/profile", &controllers.MainController{}, "get:HelloSitepoint")
	//beego.Router("/test", &controllers.MainController{}, "get:Test")
	//beego.Router("/echoService",&controllers.MainControllers{},"get:EchoService")
	//beego.Router("/echoService",&controllers.MainController{},"get:EchoService")
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.IndexController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/user/profile", &controllers.ShowMessageController{}, "get:Get")
}
