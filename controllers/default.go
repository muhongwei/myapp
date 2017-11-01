package controllers

import (
    "github.com/astaxie/beego"
    "os"
    "fmt"
    "sync"

    "log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "myapp/controllers/echo"
    
)

const(
	port=":50052"
)
var once sync.Once
type echoServer struct{}

func(s *echoServer) Echo(ctx context.Context,in *pb.EchoRequest)(*pb.EchoReply,error){
	return &pb.EchoReply{EplayMessage:in.RequestMessage},nil
}


type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	
}
func (main *MainController) HelloSitepoint() {
    // main.Data["Website"] = "My Website"
    // main.Data["Email"] = "your.email.address@example.com"
    // main.Data["EmailName"] = "Your Name"	
    // main.Data["Id"] = main.Ctx.Input.Param(":id")
    main.Data["UserAgent"] = main.Ctx.Request.UserAgent()
    // main.Data["domain"] = main.Ctx.Input.Domain()
    main.Data["IP"] = main.Ctx.Input.IP()
    // main.Data["Header"] = main.Ctx.Request.Header
    // main.Data["GetData"] = main.Ctx.Input.GetData
    // main.Data["Proxy"] = main.Ctx.Input.Proxy()
    // main.Data["Protocol"] = main.Ctx.Input.Protocol()
    // main.Data["Host"] = main.Ctx.Request.Host
    host, err := os.Hostname()
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        main.Data["HostName"] = host
    }
    

    main.TplName = "user/profile.tpl"
    
}
func (this *MainController) EchoService(){

    once.Do(func(){
        lis,err :=net.Listen("tcp",port)
        if err != nil{
            log.Fatal("failed to listen:%v",err)
        }
        s:=grpc.NewServer()
        //pb.RegisterEchoerServer(s,&echoServer{})
        pb.RegisterEchoerServer(s,&echoServer{})
        s.Serve(lis)
        this.TplName = "index.tpl"
    })


}
func (this *MainController) Test() {
    // f, h, err := c.GetFile("uploadname")
    // if err != nil {
    //     //log.Fatal("getfile err ", err)
    // }
    // defer f.Close()
	// c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	v := this.GetSession("asta")
    if v == nil {
        this.SetSession("asta", int(1))
        this.Data["num"] = 0
    } else {
        this.SetSession("asta", v.(int)+1)
        this.Data["num"] = v.(int)
    }
    //this.TplName = "index.tpl"
    this.TplName = "user/test.tpl"
}
