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
	//fmt.Println(inputs)
	user.UserName = inputs.Get("username")

	//user.UserPassword = inputs.Get("pwd")
	var password string = inputs.Get("pwd")
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	user.UserPassword=  hex.EncodeToString(cipherStr)

	user.UserIntroduction = inputs.Get("introduction")
	err := models.SaveUser(user)
	if err == nil {
		this.TplName = "user/login.html"
	} else {
		fmt.Println(err)
		this.TplName = "error.tpl"
	}
}