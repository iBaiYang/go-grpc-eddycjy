package main

import (
	"context"
	"flag"
	pb "github.com/iBaiYang/go-grpc-eddycjy/UnaryRPC/proto"
	"google.golang.org/grpc"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = DoSayHello(client)
}

func DoSayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "abcdef"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
