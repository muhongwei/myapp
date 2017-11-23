package main

//streamEchoClient.go

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	pb "myapp/streamEcho"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "192.168.34.10:10009", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func runEcho(client pb.EchoerClient) {
	stream, err := client.Echo(context.Background())
	if err != nil {
		log.Fatalf("%v.Echo(_)= _,%v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			//fmt.Println(in)
			if err != nil {
				log.Fatalf("Failed to receive a note:%v", err)
			}
			//log.Print(in.EplayMessage)
			//fmt.Print("message:")
			//log.Printf("message:%s",in.EplayMessage)
			fmt.Println(in.EplayMessage)
			if in.TimeNow != "" {
				fmt.Println(in.TimeNow)
			}

		}
	}()
	fmt.Println("输入exit退出：")

	var note pb.EchoRequest
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		note.RequestMessage = scanner.Text()
		if note.RequestMessage == "exit" {
			break

		}
		if err := stream.Send(&note); err != nil {
			log.Fatalf("failed to send a not:%v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	stream.CloseSend()
	//fmt.Println("closeSend:")
	<-waitc

}
func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dail:%v", err)
	}
	defer conn.Close()
	client := pb.NewEchoerClient(conn)
	runEcho(client)
}
