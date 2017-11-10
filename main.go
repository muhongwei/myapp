package main

import (
	"os"
	_ "myapp/routers"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"net"
	"flag"
	"time"
	"io"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "myapp/streamEcho/streamEchoService"
	pb "myapp/streamEcho"

	//"google.golang.org/grpc/reflection"

)

type sEchoServer struct{}
var (
	//port       = flag.Int("port", 10009, "The server port")
	port=":10009"
)
func (s *sEchoServer) Echo(stream pb.Echoer_EchoServer) error{

	glog.Infoln("start echo server")
	var isReturnTime = true

	go func(){
		ticker := time.NewTicker(20 * time.Second)
		for {
			time := <-ticker.C
			var rPly pb.EchoReply
			rPly.TimeNow = time.Format("2006-01-02 15:04:05")
			if isReturnTime==true{
				glog.Infoln(rPly.TimeNow)
				stream.Send(&rPly)
			}else{
				break
			}
			
		}
	}()

	for{
		in,err := stream.Recv()
		glog.Infoln(in)
		if err == io.EOF{
			glog.Infoln("clientEOF")
			isReturnTime=false
			return nil
		}
		if err != nil{
			return err
		}
		var rPly pb.EchoReply
		rPly.EplayMessage=in.RequestMessage
		if err1 := stream.Send(&rPly);err1 != nil{
			return err1
		}
	}
}
func startEchoServer(){

	//lis,err := net.Listen("tcp",fmt.Sprintf("localhost:%d",*port))
	lis,err := net.Listen("tcp",port)
	if err != nil{
		glog.Fatalln("failed to listen:%v",err)
	}
	streamEchoServer := grpc.NewServer()
	s := new(sEchoServer)
	pb.RegisterEchoerServer(streamEchoServer,s)
	streamEchoServer.Serve(lis)
}
func GetMyappName() string{
	return os.Getenv("OEM")
}
func GetMyappVersion() string{
	return os.Getenv("VER")
}

func main() {
	flag.Parse()
	defer glog.Flush()

	beego.BConfig.WebConfig.Session.SessionOn = true

	
	beego.SetStaticPath("/views","views")
	beego.SetStaticPath("/log","public")

	//注册函数
	beego.AddFuncMap("GetMyappName",GetMyappName)
	beego.AddFuncMap("GetMyappVersion",GetMyappVersion)


	//开启echo server
	glog.Infoln("open echo server")
	go startEchoServer()

    
	beego.Run()
}

