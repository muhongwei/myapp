package main

import (
	"fmt"
	_ "myapp/routers"
	"github.com/astaxie/beego"

	"log"
	"net"
	"flag"
	"time"
	"io"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "streamEcho/streamEcho"

	//"google.golang.org/grpc/reflection"

)

type sEchoServer struct{}
var (
	port       = flag.Int("port", 10009, "The server port")
)
func (s *sEchoServer) Echo(stream pb.Echoer_EchoServer) error{

	go func(){
		ticker := time.NewTicker(10 * time.Second)
		for i := 0; i < 10; i++ {
			time := <-ticker.C
			//fmt.Println(time.String())
			var rPly pb.EchoReply
			rPly.TimeNow = time.String()
			fmt.Println(rPly.TimeNow)
			stream.Send(&rPly)
		}
	}()

	for{
		in,err := stream.Recv()
		fmt.Println(in)
		if err == io.EOF{
			fmt.Println("serverEOF")
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
	flag.Parse()
	lis,err := net.Listen("tcp",fmt.Sprintf("localhost:%d",*port))
	if err != nil{
		log.Fatal("failed to listen:%v",err)
	}
	streamEchoServer := grpc.NewServer()
	s := new(sEchoServer)
	pb.RegisterEchoerServer(streamEchoServer,s)
	streamEchoServer.Serve(lis)
}

func main() {
	
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启hello server
	go startEchoServer()

	beego.Run()

	

	


	
}

