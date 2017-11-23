package controllers

import (
	"myapp/models"

	"github.com/astaxie/beego"
	"github.com/golang/glog"

	"crypto/md5"
	"encoding/hex"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {

	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("login Get():")
	glog.Infoln("userip:", index.Ctx.Input.IP())
	glog.Infoln("useragent:", index.Ctx.Request.UserAgent())

	//查看session中是否存储有用户登录信息，若没有，进入登录页面，否则进入信息显示页面
	sess := index.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		index.TplName = "user/login.html"
	} else {
		index.Redirect("/user/profile", 302)
	}

}

func (index *IndexController) Post() {
	//将页面和接口调用对应的客户端ip和UserAgent信息作为日志打印到标准控制台
	glog.Infoln("login Post():")
	glog.Infoln("userip:", index.Ctx.Input.IP())
	glog.Infoln("useragent:", index.Ctx.Request.UserAgent())

	//新建一个session存储用户信息
	sess := index.StartSession()
	//新建models.User类型变量来储存用户登录信息
	var user models.User
	inputs := index.Input()
	user.UserName = inputs.Get("username")
	var password string = inputs.Get("pwd")
	//将输入密码进行MD5加密与数据库进行对比
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	user.UserPassword = hex.EncodeToString(cipherStr)

	err := models.ValidateUser(user)
	if err == nil {

		//记录用户登录信息
		glog.Infoln("username:", user.UserName)
		glog.Infoln("userpassword:", user.UserPassword)
		glog.Infoln("login success")

		//如果登录成功，将用户名和已经通过md5加密的密码存入session
		sess.Set("username", user.UserName)
		sess.Set("userpassword", user.UserPassword)

		//查找用户介绍并存入session
		var userTemp *models.User
		userTemp = models.FindUser(user)
		sess.Set("userintroduction", userTemp.UserIntroduction)

		//进入user/profile网页
		index.Redirect("/user/profile", 302)
	} else {
		//记录用户登录信息
		glog.Infoln("username:", user.UserName)
		glog.Infoln("userpassword:", user.UserPassword)
		glog.Infoln("login failed")
		index.TplName = "error.tpl"
	}
}
