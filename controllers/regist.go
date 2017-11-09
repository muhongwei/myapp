package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myapp/models"
	"crypto/md5" 
    "encoding/hex"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.TplName = "regist.html"
}

func (this *RegistController) Post() {
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
		fmt.Println("register name:",user.UserName)
		fmt.Println("register password:",user.UserPassword)
		fmt.Println("register introduction:",user.UserIntroduction)
		fmt.Println("regist success")
		this.Redirect("/login",302)
	} else {
		fmt.Println("register name:",user.UserName)
		fmt.Println("register password:",user.UserPassword)
		fmt.Println("register introduction:",user.UserIntroduction)
		fmt.Println("regist failed")
		this.TplName = "error.tpl"
	}
}