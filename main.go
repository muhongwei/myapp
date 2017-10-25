package main

import (
	_ "myapp/routers"
	"github.com/astaxie/beego"

	"log"
	"net"
 
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "helloworld/helloworld"

)

const (
	port = ":50056"
 )
 
 type server struct {}
 
 func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:in.Name}, nil
 }
 func startServer(){
	 
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

 }

func main() {

	
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启hello server
	startServer()

	beego.Run()

	


	
}

