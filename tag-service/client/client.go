package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

var port string

func main() {
	flag.StringVar(&port, "p", "8004", "port number")
	flag.Parse()

	ctx := context.Background()

	auth := Auth{
		AppKey:    "lingze",
		AppSecret: "yuke",
	}

	SERVICE_NAME := "tag-service"
	target := fmt.Sprintf("/etcdv3://go-programming-tour/grpc/%s", SERVICE_NAME)
	opts := []grpc.DialOption{grpc.WithPerRPCCredentials(&auth)}
	clientConn, _ := GetClientConn(ctx, target, opts)
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	names := []string{"PHP", "C++", "FFFF", ""}
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: s})
			log.Printf("resp:%v\n", resp)
		}(name)
	}

	wg.Wait()
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: ""})
	log.Printf("resp:%v\n", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	config := clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	r := &naming.GRPCResolver{Client: cli}
	opts = append(opts, grpc.WithBalancer(grpc.RoundRobin(r)), grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	opts = append(opts,
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				middleware.StreamContextTimeout(),
			),
		))

	Retryinterpreter := grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithMax(2),
		grpc_retry.WithCodes(
			codes.Unknown,
			codes.Internal,
			codes.DeadlineExceeded,
		),
	)

	opts = append(opts,
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				middleware.UnaryContextTimeout(),
				Retryinterpreter,
			),
		))
	// opts = append(opts,
	// 	grpc.WithUnaryInterceptor(
	// 		grpc_middleware.ChainUnaryClient(interpreter)))

	return grpc.DialContext(ctx, target, opts...)
}

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}
