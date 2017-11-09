package controllers

import (
	"github.com/astaxie/beego"   
	"os" 
	"fmt"
)


type ShowMessageController struct {
	beego.Controller
}


func (c *ShowMessageController) Get() {
	sess := c.StartSession()
	// c.Data["UserAgent"] = sess.Get("useragent")
	// c.Data["IP"] = sess.Get("ip")
	// c.Data["HostName"] = sess.Get("hostname")
	c.Data["UserAgent"] = c.Ctx.Request.UserAgent()
	c.Data["IP"] = c.Ctx.Input.IP()
	host, err := os.Hostname()
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        c.Data["HostName"] = host
	}
	
	c.Data["UserName"] = sess.Get("username")
	c.Data["UserIntroduction"] = sess.Get("userintroduction")

	c.TplName = "user/profile.tpl"
	
}
func (c *ShowMessageController) Post(){

}