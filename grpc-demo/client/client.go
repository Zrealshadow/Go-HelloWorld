package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Lingze"})
	log.Printf("client.SayHello resp:%s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Printf("resp :%v", resp)
	}
	return nil
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	space := 2
	i := 0

	for i < len(r.Name) {
		var last int
		if i+space < len(r.Name) {
			last = i + space
		} else {
			last = len(r.Name)
		}
		_ = stream.Send(&pb.HelloRequest{Name: r.Name[i:last]})
		i = i + space
	}

	// for n := 0; n < 6; n++ {
	// 	_ = stream.Send(r)
	// }

	resp, _ := stream.CloseAndRecv()

	log.Printf("resp err : %v", resp)
	return nil
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

		log.Printf("resp err : %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}

// func SayRoute()
func main() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	// _ = SayHello(client)
	_ = SayRecord(client, &pb.HelloRequest{Name: "Lingze"})
}
