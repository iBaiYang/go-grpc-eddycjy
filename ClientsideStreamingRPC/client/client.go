package main

import (
	"context"
	"flag"
	pb "github.com/iBaiYang/go-grpc-eddycjy/ClientsideStreamingRPC/proto"
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
	r := pb.HelloRequest{Name: "abcdef"}
	_ = SayRecord(client, &r)
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("resp err: %v", resp)
	return nil
}
