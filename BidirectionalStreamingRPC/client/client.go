package main

import (
	"context"
	"flag"
	pb "github.com/iBaiYang/go-grpc-eddycjy/BidirectionalStreamingRPC/proto"
	"google.golang.org/grpc"
	"io"
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
	_ = SayRoute(client, &r)
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}
