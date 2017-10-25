package main

//echoClient.go

import(
	"log"
	"os"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "echo/echo"
)

const(
	address ="localhost:50052"
)
func main(){
	fmt.Println("输入exit退出!")
	for{
		conn,err := grpc.Dial(address,grpc.WithInsecure())
		if err != nil{
			log.Fatal("did not connect:%v",err)
		}
		defer conn.Close()
		//c :=pb.NewEchoerServer(conn)

		//c :=pb.NewEchoerClient(conn)
		c :=pb.NewEchoerClient(conn)
		
		fmt.Print("echo:")

		var message string
		if _,err := fmt.Scanf("%s\n",&message);err != nil{
			fmt.Printf("%s\n",err)
		}
		if message == "exit"{
			break
		}

		if len(os.Args) > 1{
			message = os.Args[1]
		}

		//r,err := c.Echo(context.Background(),&pb.EchoRequest{EplayMessage:message})
		r,err := c.Echo(context.Background(),&pb.EchoRequest{RequestMessage:message})
		if err != nil{
			log.Fatal("could not echo:%v",err)
		}
		fmt.Printf("%s\n",r.EplayMessage)
	}
}