package main

import (
	_ "myapp/routers"
	"github.com/astaxie/beego"

	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "echo/echo"

)

const(
	port=":50052"
)
type echoServer struct{}

func(s *echoServer) Echo(ctx context.Context,in *pb.EchoRequest)(*pb.EchoReply,error){
	return &pb.EchoReply{EplayMessage:in.RequestMessage},nil
}
 func startEchoServer(){
	 
	lis,err :=net.Listen("tcp",port)
	if err != nil{
		log.Fatal("failed to listen:%v",err)
	}
	s:=grpc.NewServer()
	pb.RegisterEchoerServer(s,&echoServer{})
	s.Serve(lis)

 }

func main() {

	
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启hello server
	startEchoServer()

	beego.Run()

	


	
}

