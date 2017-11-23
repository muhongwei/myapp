package controllers

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type ShowMessageController struct {
	beego.Controller
}

func (c *ShowMessageController) Get() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("showMessage Get():")
	glog.Infoln("userip:", c.Ctx.Input.IP())
	glog.Infoln("useragent:", c.Ctx.Request.UserAgent())

	sess := c.StartSession()
	// c.Data["UserAgent"] = sess.Get("useragent")
	// c.Data["IP"] = sess.Get("ip")
	// c.Data["HostName"] = sess.Get("hostname")
	c.Data["UserAgent"] = c.Ctx.Request.UserAgent()
	c.Data["IP"] = c.Ctx.Input.IP()
	host, err := os.Hostname()
	if err != nil {
		glog.Infof("%s", err)
	} else {
		c.Data["HostName"] = host
	}

	c.Data["UserName"] = sess.Get("username")
	c.Data["UserIntroduction"] = sess.Get("userintroduction")

	c.TplName = "user/profile.tpl"

}
func (c *ShowMessageController) Post() {

}
