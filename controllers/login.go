package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myapp/models"
	"os"

	"crypto/md5" 
    "encoding/hex" 
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {


	index.Data["UserAgent"] = index.Ctx.Request.UserAgent()
	index.Data["IP"] = index.Ctx.Input.IP()
	host, err := os.Hostname()
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        index.Data["HostName"] = host
    }

	sess := index.StartSession()
	username := sess.Get("username")
	userintroduction := sess.Get("userintroduction")
	index.Data["UserName"] = username
	index.Data["UserIntroduction"] = userintroduction
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplName = "user/login.html"
		
	} else {
		index.TplName = "user/profile.tpl"
	}

}

func (index *IndexController) Post() {

	index.Data["UserAgent"] = index.Ctx.Request.UserAgent()
	index.Data["IP"] = index.Ctx.Input.IP()
	host, err1 := os.Hostname()
    if err1 != nil {
        fmt.Printf("%s", err1)
    } else {
        index.Data["HostName"] = host
    }

	
	sess := index.StartSession()
	//var user models.User
	var user models.User
	
	inputs := index.Input()
	//fmt.Println(inputs)
	user.UserName = inputs.Get("username")
	//user.userName=
	fmt.Println(user.UserName)

	var password string = inputs.Get("pwd")
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	user.UserPassword=  hex.EncodeToString(cipherStr)
	
	fmt.Println(user.UserPassword)

	err := models.ValidateUser(user)
	if err == nil {
		sess.Set("username", user.UserName)
		fmt.Println("username:", sess.Get("username"))

		var userTemp *models.User
		userTemp = models.FindUser(user)
		sess.Set("userintroduction",userTemp.UserIntroduction)
		index.Data["UserName"] = userTemp.UserName
		index.Data["UserIntroduction"] = userTemp.UserIntroduction


		index.TplName = "user/profile.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}