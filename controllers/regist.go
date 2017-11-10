package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"myapp/models"
	"crypto/md5" 
    "encoding/hex"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("regist Get():")
	glog.Infoln("userip:",this.Ctx.Input.IP())
	glog.Infoln("useragent:",this.Ctx.Request.UserAgent())

	this.TplName = "regist.html"
}

func (this *RegistController) Post() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("regist Post():")
	glog.Infoln("userip:",this.Ctx.Input.IP())
	glog.Infoln("useragent:",this.Ctx.Request.UserAgent())

	var user models.User
	inputs := this.Input()	
	user.UserName = inputs.Get("username")	
	//将密码进行MD5加密
	var password string = inputs.Get("pwd")
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	user.UserPassword=  hex.EncodeToString(cipherStr)

	user.UserIntroduction = inputs.Get("introduction")
	//将用户信息保存到数据库，跳回登录页面。保存失败进入err页面
	err := models.SaveUser(user)
	if err == nil {
		glog.Infoln("register name:",user.UserName)
		glog.Infoln("register password:",user.UserPassword)
		glog.Infoln("register introduction:",user.UserIntroduction)
		glog.Infoln("regist success")
		this.Redirect("/login",302)
	} else {
		glog.Infoln("register name:",user.UserName)
		glog.Infoln("register password:",user.UserPassword)
		glog.Infoln("register introduction:",user.UserIntroduction)
		glog.Infoln("regist failed")
		this.TplName = "error.tpl"
	}
}