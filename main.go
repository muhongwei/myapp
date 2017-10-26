package main

import (
	"fmt"
	_ "myapp/routers"
	"github.com/astaxie/beego"

	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "echo/echo"
	"google.golang.org/grpc/reflection"

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
	//s.Serve(lis)
	reflection.Register(s)
	fmt.Println("before .serve(lis)")
	if err :=s.Serve(lis);err != nil{
		log.Fatalf("failed to server:%v",err)
	}
	fmt.Println("startover")

 }

func main() {
	
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启hello server
	//startEchoServer()

	beego.Run()

	

	


	
}

